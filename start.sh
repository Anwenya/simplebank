#!/bin/sh

set -e

# echo "run db migration"
# 通过migrate可执行文件执行数据库迁移
# 该方式已不再使用，已改为在服务启动入口执行数据库迁移
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# 获得命令的参数并执行
exec "$@"