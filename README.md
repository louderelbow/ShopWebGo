# ShopWebGo

基于 Go 语言开发的 B2C 电商平台，采用 Gin + GORM + MySQL + Redis + Elasticsearch 技术栈，涵盖前台商城与后台管理系统。

---

## 技术栈

| 类别 | 技术 | 说明 |
|------|------|------|
| 语言 | Go | |
| Web 框架 | Gin | 路由、中间件、HTML 模板渲染 |
| ORM | GORM v2 | MySQL 操作、事务支持 |
| 数据库 | MySQL 8.0 | 业务数据持久化 |
| 缓存 | Redis 7 | 商品详情、导航、分类缓存（Cache-Aside） |
| 搜索引擎 | Elasticsearch | 订单全文搜索，可配置关闭降级 MySQL LIKE |
| 认证 | JWT + Session | 前台 JWT (golang-jwt/jwt v5) 无状态鉴权，后台 Session + Cookie Store |
| 密码加密 | bcrypt | golang.org/x/crypto/bcrypt |
| 验证码 | base64Captcha | 后台图形验证码 + 短信验证码 |
| 加密 | DES (Cookie 购物车) | crypto/des + PKCS5Padding |
| 图片处理 | go_image | 商品图片缩略图生成 |
| 文件存储 | 阿里云 OSS / 本地 | 配置切换 |
| 配置管理 | go-ini | app.ini 解析（启动时加载） |

---

## 功能详解

### 前台商城（`/`）

| 模块 | 功能 |
|------|------|
| 首页 | 轮播图、分类导航、推荐商品（Redis 缓存） |
| 商品浏览 | 分类筛选、商品详情（关联商品/赠品/颜色/配件/规格参数） |
| 购物车 | 增删改、数量加减、选中切换、DES 加密 Cookie 存储（未登录可用） |
| 用户认证 | 手机号注册（短信验证码 + bcrypt 密码哈希）、登录/登出、JWT Cookie 鉴权 |
| 订单结算 | 收货地址管理、条件更新防超卖、GORM 事务原子性、Session orderSign 防重复提交 |
| 错误提示 | 下单失败时页面顶部红框显示具体原因（库存不足/地址缺失等） |
| 用户中心 | 订单列表/详情、Elasticsearch 搜索订单（Bool Query 用户隔离）、收货地址管理 |

### 后台管理系统（`/admin`）

| 模块 | 功能 |
|------|------|
| 登录 | 验证码 + Session 状态管理 |
| 管理员管理 | 增删改查、角色分配、bcrypt 密码 |
| RBAC 权限 | 四张表三级权限模型（Manager→Role→Access 自关联树→RoleAccess），Preload 一条查询渲染完整权限树 |
| 商品管理 | 增删改查、上下架、多图上传、富文本编辑、异步 goroutine + WaitGroup 生成缩略图 |
| 分类管理 | 商品分类 CRUD、父子分类层级 |
| 商品类型 | 类型与属性管理（规格参数） |
| 导航管理 | 顶部/中部导航配置、关联商品 |
| 轮播图 | 焦点图 CRUD |
| 系统设置 | 站点配置、OSS/ES 配置 |
| 缓存管理 | 一键清除 Redis 缓存 |

---

## 架构亮点

- **JWT + Session 混合认证**：前台 JWT Cookie（HttpOnly）无状态鉴权，后台 Session + Cookie Store，根据场景选方案
- **三级 RBAC 权限模型**：Manager → Role → Access（ModuleId 自关联形成模块→菜单→操作三级树）→ RoleAccess 多对多关联，SessionJudge 中间件请求级拦截，超管白名单跳过
- **订单原子性**：GORM Transaction 包裹库存条件更新 + Order + OrderItem 创建，MySQL 行级原子 UPDATE 防超卖，Session 一次性签名防重复提交
- **前端错误反馈**：下单失败时通过 Session 传递具体错误消息，前端模板渲染红色警告框（库存不足/地址缺失/数据库错误），用户无需猜测
- **ES 搜索 + MySQL 降级**：订单搜索优先 ES Bool Query（uid 隔离），ES 不可用时自动降级 MySQL LIKE
- **多级 Redis 缓存**：商品详情/导航/分类 Cache-Aside 模式（TTL 3600s），后台一键 FlushAll
- **双存储**：本地磁盘 / 阿里云 OSS 配置切换，异步 goroutine 生成多尺寸缩略图
- **优雅关闭**：signal.Notify 监听 SIGINT/SIGTERM → srv.Shutdown 等待现有请求 → 释放 DB + Redis 连接
- **安全实践**：bcrypt 密码哈希、参数化查询防 SQL 注入、ChangeStatus 白名单校验、敏感配置 .gitignore 排除

---

## 项目结构

```
ShopWebGo/
├── main.go                       # 入口：模板函数、中间件、优雅关闭
├── config/
│   ├── app.ini.example           # 配置模板（真实 app.ini 由 .gitignore 排除）
│   └── app.ini                   # 真实配置
├── controller/
│   ├── admin/                    # 后台控制器（12 个）
│   └── shopWeb/                  # 前台控制器（8 个）
├── model/                        # GORM 数据模型（20 张表）
├── router/
│   ├── AdminRouter.go            # 后台路由（SessionJudge 中间件）
│   └── DefaultRouter.go          # 前台路由（UserAuth JWT 中间件 + CORS）
├── util/
│   ├── MysqlCore.go              # MySQL 连接池
│   ├── RedisCore.go              # Redis + CacheDb 封装（Set/Get/FlushAll）
│   ├── ElasticsearchUtil.go      # ES 连接 + SearchOrderItems + IndexOrderItem
│   ├── JwtUtil.go                # JWT 生成/解析
│   ├── CookieUtil.go             # Cookie DES 加解密 + Set/Get
│   ├── CaptchaUtils.go           # 图形验证码
│   ├── DesUtil.go                # DES 加解密
│   ├── LoginUtil.go              # bcrypt 密码哈希/校验
│   ├── UploadImgUtil.go          # 图片上传 + OSS/本地双存储
│   ├── GeneratePicUtil.go        # 缩略图生成
│   └── middlewares/
│       ├── SessionJudge.go       # 后台 Session + RBAC 鉴权
│       ├── UserAuth.go           # 前台 JWT Cookie 鉴权
│       └── Cors.go              # CORS 跨域
├── templates/                    # Go HTML 模板
│   ├── admin/                    # 后台页面
│   └── shopWeb/                  # 前台页面
├── static/                       # 静态资源（CSS/JS/图片）
├── sql/
│   └── GinShop.sql               # 建表 + 示例数据
├── .gitignore
└── README.md
```

---

## 数据库表

| 表名 | 说明 |
|------|------|
| `user` | 前台用户（bcrypt 密码哈希） |
| `user_temp` | 注册验证码临时表 |
| `address` | 收货地址 |
| `order` | 订单主表 |
| `order_item` | 订单商品明细 |
| `goods` | 商品主表 |
| `goods_cate` | 商品分类 |
| `goods_type` | 商品类型 |
| `goods_type_attribute` | 商品类型属性 |
| `goods_attr` | 商品规格参数 |
| `goods_color` | 商品颜色 |
| `goods_image` | 商品图库 |
| `nav` | 导航配置 |
| `focus` | 轮播图 |
| `manager` | 后台管理员 |
| `role` | 角色 |
| `access` | 权限节点（ModuleId 自关联三级树） |
| `role_access` | 角色-权限关联 |
| `setting` | 系统配置（站点/OSS） |

---

## 快速开始

### 环境要求

- Go >= 1.21
- MySQL 8.0+
- Redis 7（可选，关闭后缓存不生效）
- Elasticsearch 9.x（可选，关闭后降级 MySQL LIKE）

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
# 编辑 app.ini，填入数据库密码等

# 启动
go mod tidy
go run main.go
```

访问地址：
- 前台商城：http://localhost:8080
- 后台管理：http://localhost:8080/admin/login（账号 admin / 密码 123456）

### 启用 Elasticsearch（可选）

```bash
# ES 启动后，在 app.ini 中配置
[elasticsearch]
enable = true
addr = https://127.0.0.1:9200
username = elastic
password = 你的ES密码
```

ES 中创建索引后，前台下单自动写入 ES，用户中心搜索走 ES。关闭 `enable = false` 后自动降级 MySQL LIKE。

---

## 配置说明

```ini
jwt_secret = 你的JWT密钥              # 前台 JWT 签名密钥
des_key = 你的DES密钥                  # 购物车 Cookie 加密密钥
session_key = 你的Session密钥          # 后台 Session 加密密钥

[mysql]
password = 你的数据库密码               # 也支持环境变量 MYSQL_PASSWORD

[redis]
redisEnable = true                     # 关闭则缓存不生效

[elasticsearch]
enable = true                          # 关闭则走 MySQL LIKE 搜索

[oss]
status = 0                             # 0=本地存储 1=阿里云 OSS
```

---

## License

MIT
