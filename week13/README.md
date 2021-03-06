## week13
对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：
微服务架构（BFF、Service、Admin、Job、Task 分模块）
API 设计（包括 API 定义、错误码规范、Error 的使用）
gRPC 的使用
Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
并发的使用（errgroup 的并行链路请求）
微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
缓存的使用优化（一致性处理、Pipeline 优化）

## 相关架构图参考
/doc

## 系统设计概述
设计一个面向B端的SAAS企业在线学习平台。目前的业务主要分课程管理、学习管理、考试管理、系统管理、报表服务、后台服务等功能。将平台微服务化，采用微服务的架构对其进行开发，先按照业务来进行微服务的划分。
## 系统项目设计概要
### 使用krota作为整个微服务的框架
### 微服务之间使用grpc进行通讯
### 客户端接入层与后台服务使用http restful风格的api进行交互
### 为了使各个服务更好的隔离，各个服务之间代码仓库进行隔离
### 服务命名规则
com.athena.(一级业务名).(二级业务名).具体业务名
### 具体规划如下
BFF服务:com.athena.bff
课程服务:com.athena.course
考试服务:com.athena.kaoshi
学习服务:com.athena.learning
系统管理服务:com.athena.system
后台服务:com.athena.admin
报表服务:com.athena.bi

##