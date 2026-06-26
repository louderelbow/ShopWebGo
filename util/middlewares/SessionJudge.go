package middlewares

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"encoding/json"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionJudge() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1、获取Url访问的地址
		pathname := strings.Split(c.Request.URL.String(), "?")[0]

		// 2、获取Session里面保存的用户信息
		session := sessions.Default(c)
		userinfo := session.Get("userinfo")
		userinfoStr, ok := userinfo.(string)

		if ok {
			var userinfoStruct []model.Manager
			err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

			// Session过期判断
			if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
				if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
					c.Redirect(302, "/admin/login")
					c.Abort() // 中断请求
					return
				}
			} else { // 用户登录成功，权限判断
				urlPath := strings.Replace(pathname, "/admin/", "", 1)

				// 不是超级管理员且访问路径不在白名单
				if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
					// 根据角色获取权限列表
					roleAccess := []model.RoleAccess{}
					util.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
					roleAccessMap := make(map[int]int)
					for _, v := range roleAccess {
						roleAccessMap[v.AccessId] = v.AccessId
					}

					// 获取当前访问的url对应的权限id
					access := model.Access{}
					util.DB.Where("url = ?", urlPath).Find(&access)

					// 判断是否有权限
					if _, ok := roleAccessMap[access.Id]; !ok {
						c.String(200, "没有权限")
						c.Abort() // 中断请求
						return
					}
				}
			}
		} else { // 没有登录
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
				c.Abort() // 中断请求
				return
			}
		}

		// 验证通过，放行
		c.Next()
	}
}

//排除权限判断的方法

func excludeAuthPath(urlPath string) bool {
	config := util.GetConfig()
	excludeAuthPathStr := config.Section("admin").Key("excludeAuthPath").String()
	if excludeAuthPathStr == "" {
		return false
	}
	excludeAuthPathSlice := strings.Split(excludeAuthPathStr, ",")
	for _, p := range excludeAuthPathSlice {
		if strings.TrimSpace(p) == urlPath {
			return true
		}
	}
	return false
}
