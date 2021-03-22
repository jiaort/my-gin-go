##项目文档
---

#### `clone`项目
```bash
git clone ssh://devtfs.cmit.local:22/DefaultCollection/Back-End%20Team/_git/my-gin-go
```

#### `copy`并修改配置文件
```bash
cd my-gin-go
cp .docker-compose/shell/config.yaml .
```

#### `go mod`使用
```bash
go mod tidy
```

#### `swagger`生成自动化API文档
##### 安装`swagger`
```bash
go get -u github.com/swaggo/swag/cmd/swag
```
###### 生成文档
```bash
swagger init
```

#### `air`热重启(自动编译运行，无需手动重启--开发环境使用)
##### 安装`air`
```bash
go get -u github.com/cosmtrek/air
```
##### 编译`air`
```bash
air -c .air.conf
```
##### 运行项目
```bash
air
```

#### 项目结构
```bash
.
├── api                                  （API）
├── config                               （配置包）
├── core                                 （核心组件）
├── docs                                  (swagger文档目录)
├── global                               （全局对象）
├── initialize                           （初始化组件）
├── logs                                 （日志归档后文件）
├── middleware                           （中间件）
├── model                                （数据库对应结构体层）
├── python                               （AI平台对象检测脚本）
├── router                               （路由层）
├── service                              （API对应函数）
├── uploads                              （文件上传）
├── utils                                （公共函数）
├── .air.conf                            （热重启配置）
├── config.yaml                          （配置文件）
├── docker-compose.yaml                  （使用docker部署项目）
├── dockerfile                           （镜像文件）
└── go.mod                               （类似于requirements）
```

#### `docker`部署项目
##### 项目部署路径
```bash
# 后端
/home/go/src/my-gin-go
# 前端
/home/web/fantasy-web
```
##### `docker-compose`一键部署并启动项目
```bash
# 使用docker-compose启动四个容器
docker-compose up
# 如果修改了某些配置选项,可以使用此命令重新打包镜像
docker-compose up --build
# 使用docker-compose 后台启动
docker-compose up -d
```

#### `Demo`地址
> [AI平台](http://47.112.199.203:9999)
