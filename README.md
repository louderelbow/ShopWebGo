# ShopWebGo

基于 Go 语言开发的电商平台，采用 Gin + GORM + Redis + Elasticsearch 技术栈，涵盖前台商城与后台管理系统。

---

## 技术栈

| 类别 | 技术                  | 说明 |
|------|---------------------|------|
| 语言 | Go 1.26             | |
| Web 框架 | Gin v1.12           | 路由、中间件、模板渲染 |
| ORM | GORM v1.31          | MySQL 操作、事务支持 |
| 数据库 | MySQL 8.0           | 业务数据持久化 |
| 缓存 | Redis 7             | 商品详情、导航、分类缓存 |
| 搜索引擎 | Elasticsearch 9.4.0 | 订单全文搜索（可选开关） |
| 认证 | JWT + Session       | 前台 JWT 无状态鉴权，后台 Session + RBAC |
| 验证码 | base64Captcha       | 后台图形验证码 + 短信验证码 |
| 加密 | DES + MD5           | Cookie 加密、密码哈希 |
| 图片处理 | go_image            | 商品图片缩放裁剪 |
| 文件存储 | 阿里云 OSS             | 可选 |
| 配置管理 | go-ini              | app.ini 解析 |

---

## 功能详解

### 前台商城（`/`）

| 模块 | 功能 |
|------|------|
| 首页 | 轮播图、分类导航、推荐商品 |
| 商品浏览 | 分类筛选、商品详情（关联商品/赠品/颜色/配件/规格参数）、Redis 缓存 |
| 购物车 | 增删改、数量加减、选中切换、DES 加密 Cookie 存储 |
| 用户认证 | 手机号注册/登录/登出、JWT 无状态鉴权 |
| 订单结算 | 收货地址管理、订单签名防重复提交、事务保护 |
| 用户中心 | 订单列表/详情、Elasticsearch 搜索订单、收货地址管理 |

### 后台管理系统（`/admin`）

| 模块 | 功能 |
|------|------|
| 登录 | 验证码 + Session 状态管理 |
| 管理员管理 | 增删改查、角色分配 |
| RBAC 权限 | 角色管理、权限节点、角色授权、页面级访问控制 |
| 商品管理 | 增删改查、上下架、多图上传、富文本编辑 |
| 分类管理 | 商品分类 CRUD、父子分类层级 |
| 商品类型 | 类型与属性管理（规格参数） |
| 导航管理 | 顶部/中部导航配置、关联商品 |
| 轮播图 | 焦点图 CRUD |
| 系统设置 | 站点配置、OSS 配置 |
| 缓存管理 | 一键清除 Redis 缓存 |

---

## 架构亮点

- **JWT + Session 混合认证**：前台采用 JWT 无状态鉴权，后台采用 Session + RBAC 权限模型，兼顾扩缩容与精细权限控制
- **Elasticsearch 搜索 + MySQL 降级**：订单搜索优先 ES，故障自动降级 MySQL LIKE，控制台输出当前使用的搜索方式
- **多级 Redis 缓存**：商品详情、导航、分类数据缓存（TTL 1 小时），首次查库后续命中缓存
- **事务保护**：下单操作使用 `DB.Transaction` 包裹订单与订单项写入，失败全回滚
- **SQL 注入防护**：全项目使用 GORM 参数化查询（`?` 占位符）
- **防重复提交**：订单页随机签名存入 Session，提交校验后销毁
- **DES Cookie 加密**：购物车数据序列化后 DES 加密存储
- **优雅关闭**：监听 SIGINT/SIGTERM，平滑关闭 HTTP 服务后释放数据库与 Redis 连接
- **暗码安全**：数据库密码、OSS 密钥等敏感信息支持环境变量读取，配置文件通过 `.gitignore` 排除

---

## 项目结构

```
ShopWebGo/
├── main.go                     # 入口：路由注册、模板函数、中间件、优雅关闭
├── go.mod / go.sum             # Go 模块依赖
├── .air.toml                   # Air 热重载配置
├── config/
│   ├── app.ini.example         # 配置文件模板
│   └── app.ini                 # 真实配置（Git 忽略）
├── controller/
│   ├── admin/                  # 后台控制器（12个）
│   │   ├── LoginController.go      # 登录/登出/验证码
│   │   ├── ManagerController.go    # 管理员 CRUD
│   │   ├── RoleController.go       # 角色管理 + 授权
│   │   ├── AccessController.go     # 权限节点管理
│   │   ├── GoodsController.go      # 商品管理 + 图片/富文本上传
│   │   ├── GoodsCateController.go  # 商品分类
│   │   ├── GoodsTypeController.go  # 商品类型
│   │   ├── GoodsTypeAttributeController.go  # 商品属性
│   │   ├── FocusController.go      # 轮播图
│   │   ├── NavController.go        # 导航管理
│   │   ├── SettingController.go    # 系统设置
│   │   ├── MainController.go       # 首页/状态切换/缓存清除
│   │   └── BaseController.go       # 管理端基础渲染
│   └── shopWeb/                # 前台控制器（8个）
│       ├── DefaultController.go    # 首页
│       ├── ProductController.go    # 商品列表/详情/图库（Redis 缓存）
│       ├── CartController.go       # 购物车 CRUD
│       ├── PassController.go       # 登录/注册（JWT 签发）
│       ├── CheckOutController.go   # 订单结算（事务 + ES 同步 + 模拟支付）
│       ├── AddressController.go    # 收货地址
│       ├── UserController.go       # 用户中心/订单查询（ES 搜索优先）
│       └── BaseController.go       # 前台公共渲染（导航/分类/用户状态）
├── model/                      # 数据模型（20 张表）
├── router/
│   ├── AdminRouter.go          # 后台路由（Session 鉴权中间件）
│   └── DefaultRouter.go        # 前台路由（JWT 鉴权 + CORS）
├── util/
│   ├── MysqlCore.go            # MySQL 连接（环境变量优先）
│   ├── RedisCore.go            # Redis 连接 + CacheDb 封装
│   ├── ElasticsearchUtil.go    # ES 连接 + SearchOrderItems + IndexOrderItem
│   ├── JwtUtil.go              # JWT 生成/解析/获取用户
│   ├── CookieUtil.go           # Cookie DES 加解密
│   ├── CaptchaUtils.go         # 图形验证码生成校验
│   ├── DesUtil.go              # DES 加解密算法
│   ├── GetGoodsUtil.go         # 商品查询工具
│   ├── UploadImgUtil.go        # 图片上传 + OSS
│   ├── GeneratePicUtil.go      # 缩略图生成
│   ├── FormatAttrUtil.go       # 规格属性格式化
│   ├── TransformUtil.go        # 类型转换（Int/String/Md5 等）
│   ├── TimeUtils.go            # 时间处理
│   ├── GetOrderIdUtil.go       # 订单号生成
│   ├── LoginUtil.go            # 站点配置读取
│   └── middlewares/
│       ├── SessionJudge.go     # 后台 Session + RBAC 鉴权
│       ├── UserAuth.go         # 前台 JWT 鉴权
│       └── Cors.go            # CORS 跨域
├── templates/                  # HTML 模板（Go template）
│   ├── admin/                  # 后台页面
│   └── shopWeb/                # 前台页面
├── static/                     # 静态资源
│   ├── admin/                  # Bootstrap + 后台 CSS/JS
│   ├── shopWeb/                # 前台 CSS/JS/图片
│   └── diyUpload/              # 文件上传组件
├── sql/
│   └── GinShop.sql             # 建表语句 + 示例数据
├── .gitignore
└── README.md
```

---

## 数据库表

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
| `user_temp` | 注册验证临时表 |
| `address` | 收货地址 |
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
- Redis 7（可选）
- Elasticsearch 9.x（可选）

### 安装与运行

```bash
git clone https://github.com/louderelbow/ShopWebGo.git
cd ShopWebGo

# 创建数据库
mysql -u root -p -e "CREATE DATABASE ginshop CHARACTER SET utf8mb4;"

# 导入数据
mysql -u root -p ginshop < sql/GinShop.sql

# 复制配置
cp config/app.ini.example config/app.ini
# 编辑 app.ini，修改数据库密码等

# 安装依赖并启动
go mod tidy
go run main.go
```

访问地址：
- 前台商城：http://localhost:8080
- 后台管理：http://localhost:8080/admin/login

### 启用 Elasticsearch（可选）

ES 安装并启动后，在 `app.ini` 配置：

```ini
[elasticsearch]
enable = true
addr = https://127.0.0.1:9200
username = elastic
password = 你的ES密码
```

在 ES 中创建索引：

```json
PUT /order_items
{
  "settings": { "number_of_replicas": 0 },
  "mappings": {
    "properties": {
      "order_id":       { "type": "integer" },
      "uid":            { "type": "integer" },
      "product_title":  { "type": "text" },
      "product_price":  { "type": "float" }
    }
  }
}
```

之后在前台下单会自动写入 ES，用户中心搜索自动走 ES。关闭 `enable = false` 则自动降级为 MySQL LIKE。

---

## 配置说明

```ini
jwt_secret = 你的JWT密钥           # 前台 JWT 签名密钥
des_key = 你的DES密钥               # 购物车 Cookie 加密密钥
session_key = 你的Session密钥       # 后台 Session 加密密钥

[mysql]
password = 你的数据库密码            # 也支持环境变量 MYSQL_PASSWORD

[redis]
redisEnable = true                  # 关闭则不用 Redis

[elasticsearch]
enable = true                       # 关闭则走 MySQL LIKE 搜索
```

---

## 后台默认管理员

| 账号 | 密码 |
|------|------|
| admin | 123456 |

---

## License

MIT
