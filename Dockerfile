# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
# 由于网络原因不使用这种方式 直接在下面从主机拷贝migrate
# 该方式已不再使用，已改为在服务启动入口执行数据库迁移
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
ENV GOPROXY=https://goproxy.io
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
# 该方式已不再使用，已改为在服务启动入口执行数据库迁移
# COPY migrate .
COPY db/migration ./db/migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
# ENTRYPOINT 和 CMD 同时存在时,将等价于 [ENTRYPOINT CMD], cmd的内容将作为参数
ENTRYPOINT [ "/app/start.sh" ]