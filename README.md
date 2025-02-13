### 项目说明
#### 项目介绍
本项目是参照 https://github.com/a54552239/pearProjectApi 来进行构建的，
原项目的后端是用php和java来写的。本项目参照其，使用go来进行重构

前端仓库：https://github.com/WildWeber-xiaocheng/ms_project_front.git

工程目录设计参考：https://github.com/golang-standards/project-layout

目录名称含义：
* cmd：可执行文件，可能有多个main文件
* internal：内部代码，不希望外部访问
* pkg：公开代码，外部可以访问
* config/configs/etc: 配置文件
* scripts：脚本
* docs：文档
* third_party: 三方辅助工具
* bin：编译的二进制文件
* build：持续集成相关
* deploy：部署相关
* test：测试文件
* api：开放的api接口
* init：初始化函数

数据库表在mysql_init文件中

#### 项目使用技术
* go版本：1.23.2
* 框架：gin
* grpc
* mysql
* redis
* 日志：zap
* 配置文件读取：viper
* 服务发现：etcd



