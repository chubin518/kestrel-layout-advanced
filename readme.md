# kestrel-layout-advanced

基于 Golang 的应用脚手架

```
kestrel-layout-advanced
.
├── buildinfo
│   └── buildinfo.go
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── conf
│   ├── dev.yaml
│   ├── prod.yaml
│   ├── stage.yaml
│   └── test.yaml
├── dist
│   ├── app_brain
│   └── conf
│       └── dev.yaml
├── go.mod
├── go.sum
├── internal
│   ├── model
│   │   ├── log.go
│   │   └── props.go
│   ├── routes
│   │   ├── handler
│   │   │   └── log.go
│   │   └── routes.go
│   └── service
│       └── log.go
├── pkg
│   ├── config
│   │   ├── config.go
│   │   ├── global.go
│   │   ├── options.go
│   │   └── viper.go
│   ├── errors
│   │   └── errors.go
│   ├── graceful
│   │   ├── graceful.go
│   │   ├── hook.go
│   │   ├── routes.go
│   │   ├── timeout.go
│   │   └── web_graceful.go
│   ├── logging
│   │   ├── global.go
│   │   ├── level.go
│   │   ├── logging.go
│   │   ├── options.go
│   │   ├── utils.go
│   │   └── zap.go
│   ├── middleware
│   │   ├── logging.go
│   │   ├── recovery.go
│   │   └── requestid.go
│   ├── page
│   │   └── page.go
│   ├── result
│   │   └── result.go
│   └── utils
│       ├── compress.go
│       ├── file.go
│       ├── http.go
│       ├── number.go
│       ├── time.go
│       └── url.go
├── readme.md
├── scripts
│   └── build.sh
├── test
│   ├── config_test.go
│   ├── graceful_test.go
│   └── logging_test.go
└── webui
    ├── dist
    │   └── index.html
    └── webui.go
```

### 使用 nunu 创建项目

```
# nunu官网
https://github.com/go-nunu/nunu/

# 安装nunu

go install github.com/go-nunu/nunu@latest

# 创建项目
nunu new projectName -r https://github.com/chubin518/kestrel-layout-advanced.git

```

### 启动项目

```
# 构建webui
cd webui
yarn install
yarn run build

# 回到根目录
cd ..
go run ./cmd

# build项目
./scripts/build.sh

# 启动build项目
./dist/app

```

### CGO跨平台构建

```
# 安装zig
brew install zig

# 查看对应平台C/C++编译器
zig targets | grep gnu


# musl-cross
https://github.com/FiloSottile/homebrew-musl-cross

# 安装linux编译工具链
brew install FiloSottile/musl-cross/musl-cross --with-aarch64

# 查看对应平台C/C++编译器
/usr/local/bin

# 默认其他平台编译器
brew info musl-cross

# 目标服务器安装musl依赖

# ubuntu
apt-get update -y
apt-get install -y musl

# centos
wget https://copr.fedorainfracloud.org/coprs/ngompa/musl-libc/repo/epel-7/ngompa-musl-libc-epel-7.repo -O /etc/yum.repos.d/ngompa-musl-libc-epel-7.repo
yum install -y musl-libc-static


# 参考
https://www.baifachuan.com/posts/4862a3b1.html
https://www.sakishum.com/2021/11/29/Golang-%E4%BA%A4%E5%8F%89%E7%BC%96%E8%AF%91%E6%8A%A5%E9%94%99-XX-is-invalid-in-C99/#/%E7%8E%AF%E5%A2%83
```
