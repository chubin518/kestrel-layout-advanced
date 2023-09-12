#!/bin/bash

# 应用版本
APP_VERSION="0.0.1-beta"

# 应用名称
APP_NAME="app"

# 应用环境 dev/test/stage/prod
APP_ENV="dev"

# 目标系统 linux/windows/darwin
GO_OS="linux"

# 目标指令集 amd64/arm64
GO_ARCH="arm64"

echo "building webui"

cd ./webui
rm -rf ./dist
yarn install
yarn run build
cd ..

echo "Deleting build artifacts"

BUILD_DIR="dist"

rm -rf $BUILD_DIR
mkdir -p $BUILD_DIR/conf
cp ./conf/$APP_ENV.yaml $BUILD_DIR/conf/$APP_ENV.yaml

echo "Installing dependencies"

go mod tidy

# cross platform build args
GO_CC="`go env CC`"
GO_CXX="`go env CXX`"
GO_AR="`go env AR`"

if [ "$GO_OS" != "`go env GOOS`" ] || [ "$GO_ARCH" != "`go env GOARCH`" ]; then
    echo "Cross platform building ${APP_NAME} ${APP_VERSION} ${GO_OS}/${GO_ARCH}..."
    if [ "$GO_OS" == "linux" ] && [ "$GO_ARCH" == "arm64" ]; then
        GO_CC="zig cc -target aarch64-linux-gnu -isystem /usr/include -L/usr/lib/aarch64-linux-gnu"
        GO_CXX="zig c++ -target aarch64-linux-gnu -isystem /usr/include -L/usr/lib/aarch64-linux-gnu"
        GO_AR="zig ar -target aarch64-linux-gnu -isystem /usr/include -L/usr/lib/aarch64-linux-gnu"
    elif [ "$GO_OS" == "linux" ] && [ "$GO_ARCH" == "amd64" ];then
        GO_CC="zig cc -target x86_64-linux-gnu -isystem /usr/include -L/usr/lib/x86_64-linux-gnu"
        GO_CXX="zig c++ -target x86_64-linux-gnu -isystem /usr/include -L/usr/lib/x86_64-linux-gnu"
        GO_AR="zig ar -target x86_64-linux-gnu -isystem /usr/include -L/usr/lib/x86_64-linux-gnu"
    fi
else
    echo "Building ${APP_NAME} ${APP_VERSION} ${GO_OS}/${GO_ARCH}..."
fi

# 输出文件
OUT_PUT=$BUILD_DIR/${APP_NAME}

# build info
BUILD_FLAGS="-s -w -linkmode external -extldflags -static -X github.com/chubin518/kestrel-layout-advanced/buildinfo.Version=$APP_VERSION \
-X github.com/chubin518/kestrel-layout-advanced/buildinfo.Name=$APP_NAME \
-X github.com/chubin518/kestrel-layout-advanced/buildinfo.Environment=$APP_ENV \
-X 'github.com/chubin518/kestrel-layout-advanced/buildinfo.BuildTime=`date "+%Y-%m-%d %H:%M:%S"`' \
-X 'github.com/chubin518/kestrel-layout-advanced/buildinfo.BuildVersion=`go version`'"

# build command
CGO_ENABLED=1 GOOS=$GO_OS GOARCH=$GO_ARCH CC=$GO_CC CXX=$GO_CXX AR=$GO_AR \
go build -o $OUT_PUT -ldflags "$BUILD_FLAGS" ./cmd \

echo "Build successfully"

chmod a+x $OUT_PUT

echo "Run App ./$OUT_PUT"
