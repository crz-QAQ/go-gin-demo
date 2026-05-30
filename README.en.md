# go-gin-demo

### Introduction

Go + Gin + GORM + MySQL + Redis integration development example, providing a standard Web project skeleton to implement interface services, database operations, Redis caching, and other common features.

This is a RESTful API backend project developed using the Go language and Gin framework, integrating GORM as the ORM, MySQL as the database, Redis as the cache, and including the integration of the message queue RabbitMQ.

### Software Architecture
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
### technology stack

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM framework
- **MySQL** - Relational database
- **Redis** - Cache database
- **RabbitMQ** - message queue

### Features

- User Registration and Login (Token Authentication)
- User Information Management (CRUD, Soft Delete)
- Personal Details Management
- Message Release and Review
- User Check-in System
- Redis Cache Operations
- Unified response format
- Role Permission Control

### Installation Tutorial

1. Make sure the Go 1.18+ environment is installed
2. Clone the project and install dependencies：
``` bash
go mod download
```
3. Configure MySQL database (create database and import SQL file)
4. Configure Redis service
5. Modify the database connection information in the configuration file

6. Run the project：
``` bash
go run main.go
```

### Instructions for Use
1. The interface service runs by default on `http://localhost:8080`

2. Account-related interfaces：
   - POST `/api/account/register` - Register
   - POST `/api/account/login` - Login
   - GET `/api/account/personal_msg` - Get personal information
   - POST `/api/account/logout` - Log out
   - PUT `/api/account/update_nickname` - Modify Nickname
   - DELETE `/api/account/delete` - Delete Account
   - PUT `/api/account/restore` - Recover account

3. Message-related interfaces：
   - POST `/api/message/create` - Create message
   - GET `/api/message/list` - Get message list
   - GET `/api/message/detail` - Message Details
   - PUT `/api/message/audit` - Review message
   - DELETE `/api/message/delete` - Delete message
4. Check-in related interfaces：
   - POST `/api/sign/sign` - User check-in
   - GET `/api/sign/user_list` - User Sign-in Records
   - GET `/api/sign/admin_list` - Administrator Sign-in Records
5. Redis Operation Interface：
   - POST `/api/redis/set` - Set cache
   - GET `/api/redis/get` - Get cache
   - DELETE `/api/redis/del` - Delete cache

### Contribute
1. Fork this repository
2. Create new Feat_xxx branch
3. Submit code
4. New Pull Request
