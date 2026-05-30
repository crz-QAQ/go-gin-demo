

# go-gin-demo

### 介绍

Go + Gin + GORM + MySQL + Redis 整合开发示例，提供标准 Web 项目骨架实现接口服务、数据库操作、Redis 缓存等常用功能。

这是一个基于 Go 语言和 Gin 框架开发的 RESTful API 后端项目，集成了 GORM 作为 ORM、MySQL 作为数据库、Redis 作为缓存，并包含消息队列 RabbitMQ 的集成。

### 软件架构

```
go-gin-demo/
├── main.go                 # 程序入口
├── api/                   # API 处理器层
│   ├── account_api/       # 账号相关接口
│   ├── message_api/      # 消息相关接口
│   ├── redis_api/        # Redis 操作接口
│   ├── sign_api/         # 签到相关接口
│   └── test_user_api/    # 用户测试接口
├── dao/                   # 数据访问层
│   ├── account_dao.go
│   ├── message_dao.go
│   ├── sign_dao.go
│   └── user_dao.go
├── model/                 # 数据模型
│   ├── Account.go        # 账号模型
│   ├── Log.go            # 日志模型
│   ├── Message.go       # 消息模型
│   ├── OperateLog.go    # 操作日志模型
│   ├── Sign.go          # 签到模型
│   ├── User.go          # 用户模型
│   └── migrate.go       # 数据库迁移
├── service/               # 业务逻辑层
│   ├── account_service.go
│   ├── message_service.go
│   ├── redis_service.go
│   ├── sign_service.go
│   ├── token_service.go
│   └── user_service.go
├── middleware/            # 中间件
│   ├── auth.go           # 登录认证
│   └── role.go           # 角色权限
├── pkg/                  # 公共包
│   ├── base/             # 基础工具
│   ├── db/               # 数据库连接
│   ├── rabbitmq/        # RabbitMQ 消息队列
│   ├── redis/           # Redis 客户端
│   └── response/        # 统一响应
├── router/               # 路由配置
└── vo/                  # 视图对象
```

### 技术栈

- **Go** - 编程语言
- **Gin** - Web 框架
- **GORM** - ORM 框架
- **MySQL** - 关系型数据库
- **Redis** - 缓存数据库
- **RabbitMQ** - 消息队列

### 功能特性

- 用户注册登录（Token 认证）
- 用户信息管理（CRUD、软删除）
- 个人详情管理
- 消息发布与审核
- 用户签到系统
- Redis 缓存操作
- 统一响应格式
- 角色权限控制

### 安装教程

1. 确保已安装 Go 1.18+ 环境

2. 克隆项目并安装依赖：
   ```bash
   go mod download
   ```

3. 配置 MySQL 数据库（创建 database 并导入 SQL 文件）

4. 配置 Redis 服务

5. 修改配置文件中的数据库连接信息

6. 运行项目：
   ```bash
   go run main.go
   ```

### 使用说明

1. 接口服务默认运行在 `http://localhost:8080`

2. 账号相关接口：
   - POST `/api/account/register` - 注册
   - POST `/api/account/login` - 登录
   - GET `/api/account/personal_msg` - 获取个人信息
   - POST `/api/account/logout` - 登出
   - PUT `/api/account/update_nickname` - 修改昵称
   - DELETE `/api/account/delete` - 删除账号
   - PUT `/api/account/restore` - 恢复账号

3. 消息相关接口：
   - POST `/api/message/create` - 创建消息
   - GET `/api/message/list` - 获取消息列表
   - GET `/api/message/detail` - 消息详情
   - PUT `/api/message/audit` - 审核消息
   - DELETE `/api/message/delete` - 删除消息

4. 签到相关接口：
   - POST `/api/sign/sign` - 用户签到
   - GET `/api/sign/user_list` - 用户签到记录
   - GET `/api/sign/admin_list` - 管理员签到记录

5. Redis 操作接口：
   - POST `/api/redis/set` - 设置缓存
   - GET `/api/redis/get` - 获取缓存
   - DELETE `/api/redis/del` - 删除缓存

### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request