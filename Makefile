postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	dokcer exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

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

.PHONY: postgres createdb dropdb migrateup migratedown sqlc 
