DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres16 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	dokcer exec -it postgres16 dropdb simple_bank

# 依次向上迁移
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
# 依次向下迁移
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
# 向上迁移一次
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1
# 向下迁移一次
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

# 测试通过环境变量改变端口后会覆盖配置文件
server-port:
	SERVER_ADDRESS=0.0.0.0:7778 go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go com.wlq/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock server-port db_docs db_schema
