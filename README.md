# go-gin-demo
Go + Gin + GORM + MySQL + Redis + RabbitMQ 标准化后端项目骨架

## 介绍
基于 Go 语言生态的一站式 Web 开发脚手架，集成常用中间件与工具库，提供清晰的项目分层结构，可快速用于企业级接口服务、管理后台、微服务模块开发。

### 技术栈
1. 框架：Gin
2. ORM：GORM
3. 数据库：MySQL
4. 缓存：Redis
5. 消息队列：RabbitMQ
6. 结构：标准 MVC 分层（api/dao/model/service/router）


### 项目结构

go-gin-demo/
├── api/            # 接口层（请求入口、参数校验）
├── dao/            # 数据访问层（数据库操作）
├── middleware/     # 中间件（鉴权、日志、跨域等）
├── model/          # 数据模型（结构体、表结构）
├── pkg/            # 公共工具包（配置、MQ、工具函数）
├── router/         # 路由注册
├── service/        # 业务逻辑层
├── vo/             # 视图对象（请求/响应结构体）
├── go.mod          # Go 模块依赖
├── go.sum
└── main.go         # 项目入口

###  项目内置功能
- 统一接口返回格式封装
- 全局异常错误捕获处理
- 路由分组管理
- MySQL 基础增删改查封装
- Redis 缓存常用操作
- RabbitMQ 消息发送与消费示例
- 请求入参校验
- 跨域请求处理
- 自定义日志输出

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

