#!/bin/sh

set -e

echo "run db migration"
# 执行数据库迁移
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# 获得命令的参数并执行
exec "$@"