# ShopWebGo

基于 Go 语言开发的电商平台，采用 Gin + GORM + Redis 技术栈，涵盖前台商城与后台管理系统，支持商品展示、购物车、订单结算、RBAC 权限控制以及 Redis 缓存优化。

---

## 技术栈

| 类别 | 技术 | 说明 |
|------|------|------|
| 语言 | Go 1.26 | |
| Web 框架 | Gin v1.12 | HTTP 路由、中间件、模板渲染 |
| ORM | GORM v1.31 | MySQL 数据库操作、事务支持 |
| 数据库 | MySQL 8.0 | 业务数据持久化 |
| 缓存 | Redis 7 | 商品详情、导航、分类数据缓存 |
| Session | gin-contrib/sessions | Cookie 存储的 Session 管理 |
| 验证码 | base64Captcha | 后台登录图形验证码 |
| 加密 | DES | Cookie 数据传输加密 |
| 图片处理 | go_image | 商品图片缩放与裁剪 |
| 配置管理 | go-ini | app.ini 配置文件解析 |
| 阿里云 OSS | aliyun-oss-go-sdk | 图片对象存储（可选） |
| 二维码 | go-qrcode | 二维码生成 |

---

## 功能详解

### 前台商城（`/`）

| 模块 | 功能 |
|------|------|
| 首页 | 轮播图、商品分类导航、推荐商品展示 |
| 商品浏览 | 分类筛选、商品搜索、商品详情（关联商品、赠品、颜色、配件） |
| 购物车 | 添加/删除商品、数量加减、选中切换、DES 加密 Cookie 存储 |
| 用户认证 | 手机号注册、登录/登出、Session + Cookie 鉴权 |
| 订单结算 | 收货地址选择、订单签名防重复提交、事务保护写入 |
| 用户中心 | 订单列表、订单详情、收货地址管理 |

### 后台管理系统（`/admin`）

| 模块 | 功能 |
|------|------|
| 管理员 | 登录验证码、Session 状态管理 |
| 管理员管理 | 增删改查、角色分配 |
| RBAC 权限 | 角色管理、权限节点管理、角色授权、页面级访问控制 |
| 商品管理 | 商品增删改查、上下架、多图上传、富文本编辑 |
| 分类管理 | 商品分类增删改查、父子分类层级 |
| 商品类型 | 类型与属性管理（规格参数） |
| 导航管理 | 顶部/中部导航配置、关联商品 |
| 轮播图 | 焦点图增删改查 |
| 系统设置 | 站点配置、OSS 配置 |
| 缓存管理 | 一键清除 Redis 缓存 |

---

## 架构亮点

- **多级缓存策略**：商品详情使用 Redis 缓存，首次查询写入缓存（TTL 1小时），后续命中零数据库查询
- **事务保护**：下单操作使用 `DB.Transaction` 包裹订单与订单商品写入，保证数据一致性
- **SQL 注入防护**：全部查询使用 GORM 参数化查询（`?` 占位符），杜绝字符串拼接
- **RBAC 权限模型**：角色 → 权限节点 → URL 路由的三级权限控制，超级管理员白名单机制
- **防重复提交**：订单页生成随机签名存入 Session，提交时校验并销毁，防止刷新重复下单
- **DES Cookie 加密**：购物车数据序列化后 DES 加密存储，密钥通过配置文件管理
- **优雅关闭**：监听 SIGINT/SIGTERM 信号，平滑关闭 HTTP 服务后释放数据库与 Redis 连接
- **中间件鉴权**：后台 Session 鉴权 + 前台用户 Cookie 鉴权，CORS 跨域支持

---

## 项目结构

```
ShopWebGo/
├── main.go                     # 入口：路由注册、中间件、优雅关闭
├── go.mod / go.sum             # Go 模块依赖
├── .air.toml                   # 热重载配置
├── config/
│   ├── app.ini.example         # 配置文件模板（需复制为 app.ini）
│   └── app.ini                 # 真实配置（Git 忽略，含密码）
├── controller/
│   ├── admin/                  # 后台控制器（12个）
│   │   ├── LoginController.go      # 登录/登出/验证码
│   │   ├── ManagerController.go    # 管理员 CRUD
│   │   ├── RoleController.go       # 角色管理 + 授权
│   │   ├── AccessController.go     # 权限节点管理
│   │   ├── GoodsController.go      # 商品管理 + 图片上传
│   │   ├── GoodsCateController.go  # 商品分类
│   │   ├── GoodsTypeController.go  # 商品类型
│   │   ├── GoodsTypeAttributeController.go  # 商品属性
│   │   ├── FocusController.go      # 轮播图
│   │   ├── NavController.go        # 导航管理
│   │   ├── SettingController.go    # 系统设置
│   │   ├── MainController.go       # 后台首页/状态切换/缓存清除
│   │   └── BaseController.go       # 后台基础渲染
│   ├── shopWeb/                # 前台控制器（8个）
│   │   ├── DefaultController.go    # 首页
│   │   ├── ProductController.go    # 商品列表/详情/图库
│   │   ├── CartController.go       # 购物车 CRUD
│   │   ├── PassController.go       # 登录/注册
│   │   ├── CheckOutController.go   # 订单结算
│   │   ├── AddressController.go    # 收货地址
│   │   ├── UserController.go       # 用户中心/订单查询
│   │   └── BaseController.go       # 前台公共渲染（导航/分类/用户状态）
│   └── util/
│       └── UnknownController.go    # 工具测试（缩略图/二维码）
├── model/                      # 数据模型（20张表）
├── router/
│   ├── AdminRouter.go          # 后台路由（/admin 分组，Session 鉴权）
│   └── DefaultRouter.go        # 前台路由（/ 分组，用户鉴权中间件）
├── util/
│   ├── MysqlCore.go            # MySQL 连接初始化（环境变量优先）
│   ├── RedisCore.go            # Redis 连接 + 缓存封装
│   ├── CookieUtil.go           # Cookie 加密/解密/读写
│   ├── CaptchaUtils.go         # 图形验证码生成与校验
│   ├── DesUtil.go              # DES 加解密
│   ├── GetGoodsUtil.go         # 商品查询工具
│   ├── UploadImgUtil.go        # 图片上传处理
│   ├── GeneratePicUtil.go      # 缩略图生成
│   ├── FormatAttrUtil.go       # 属性格式化
│   ├── TransformUtil.go        # 类型转换（Int/String/Md5 等）
│   ├── TimeUtils.go            # 时间处理
│   ├── LoginUtil.go            # 站点配置读取
│   └── middlewares/
│       ├── SessionJudge.go     # 后台 Session 鉴权 + RBAC 权限校验
│       ├── UserAuth.go         # 前台用户登录鉴权
│       └── Cors.go            # CORS 跨域中间件
├── templates/                  # HTML 模板（Go template）
│   ├── admin/                  # 后台页面模板
│   └── shopWeb/                # 前台页面模板
├── static/                     # 静态资源
│   ├── admin/                  # Bootstrap + 后台 CSS/JS
│   ├── shopWeb/                # 前台 CSS/JS/图片
│   └── diyUpload/              # 文件上传组件
├── sql/
│   └── GinShop.sql             # 数据库建表 + 示例数据
├── .gitignore                  # Git 忽略规则
└── README.md
```

---

## 数据库表一览

| 表名 | 说明 |
|------|------|
| `goods` | 商品主表 |
| `goods_cate` | 商品分类 |
| `goods_type` | 商品类型 |
| `goods_type_attribute` | 商品类型属性 |
| `goods_attr` | 商品规格参数 |
| `goods_color` | 商品颜色 |
| `goods_image` | 商品图库 |
| `nav` | 导航配置 |
| `focus` | 轮播图 |
| `user` | 前台用户 |
| `user_temp` | 用户临时（注册验证） |
| `address` | 收货地址 |
| `cart` | 购物车（Cookie 为主） |
| `order` | 订单主表 |
| `order_item` | 订单商品明细 |
| `manager` | 后台管理员 |
| `role` | 角色 |
| `access` | 权限节点 |
| `role_access` | 角色-权限关联 |
| `setting` | 系统配置 |

---

## 快速开始

### 环境要求

- Go 1.26+
- MySQL 8.0+
- Redis 7（可选，配置 `redisEnable = false` 可关闭）

### 安装与运行

```bash
# 1. 克隆项目
git clone https://github.com/你的用户名/ShopWebGo.git
cd ShopWebGo

# 2. 创建数据库
mysql -u root -p -e "CREATE DATABASE ginshop CHARACTER SET utf8mb4;"

# 3. 导入表结构和数据
mysql -u root -p ginshop < sql/GinShop.sql

# 4. 复制并修改配置文件
cp config/app.ini.example config/app.ini
# 编辑 app.ini，填入你的数据库密码和 Redis 配置

# 5. 安装依赖
go mod tidy

# 6. 启动服务
go run main.go
```

访问地址：
- 前台商城：http://localhost:8080
- 后台管理：http://localhost:8080/admin/login

---

## 配置说明（config/app.ini）

```ini
[mysql]
ip       = 127.0.0.1       # 数据库地址
port     = 3306            # 数据库端口
user     = root            # 数据库用户
password = 你的密码         # 数据库密码（支持环境变量 MYSQL_PASSWORD）
database = ginshop         # 数据库名

[redis]
ip   = 127.0.0.1           # Redis 地址
port = 6379                # Redis 端口
redisEnable = true         # 是否启用 Redis 缓存
```

> 数据库密码优先从环境变量 `MYSQL_PASSWORD` 读取，未设置时回退到配置文件。

---

## 后台默认管理员

| 账号 | 密码 |
|------|------|
| admin | 123456 |

> 密码使用 MD5 存储，可在 `manager` 表中修改。

---

## License

MIT
