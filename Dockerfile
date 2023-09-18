
# build stage
FROM golang:1.21.1-alpine AS builder
ARG ENV="dev"
ENV CGO_ENABLED=1 \
    GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
WORKDIR /app
COPY . .
RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk --no-cache add \
    ca-certificates \
    git \
    gcc \
    g++ \
    pkgconfig \
    zeromq \
    zeromq-dev \
    file \
    curl \
    && mkdir -p ./dist/conf \
    && cp ./conf/${ENV}.yaml ./dist/conf/${ENV}.yaml
RUN go mod tidy \
    && go build -ldflags '-s -w -linkmode external -extldflags "-static"' -o ./dist/app_brain ./cmd
#final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dist /app
RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk --no-cache add \
    ca-certificates \
    git \
    gcc \
    g++ \
    pkgconfig \
    zeromq \
    zeromq-dev \
    file \
    curl
VOLUME [ "./app" ]
EXPOSE 8080
CMD [ "./app_brain" ]