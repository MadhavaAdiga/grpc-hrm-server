generate:
	protoc -I protos/ protos/*.proto --go_out=protos/
	protoc -I protos/ protos/*.proto --go-grpc_out=protos/

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root hrm_db

dropdb:
	docker exec -it postgres13 dropdb hrm_db	

mock:
	mockgen -package mockdb  -destination db/mock/store.go  github.com/MadhavaAdiga/grpc-hrm-server/db Store	

migrate-create:
	migrate create -ext sql -dir db/migration -seq init_schema	

migrate-up-all:
	migrate -path db/migration -database "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable" -verbose up	

migrate-down-all:
	migrate -path db/migration -database "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable" -verbose down	

# migrate up last one 
migrate-up:
	migrate -path db/migration -database "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable" -verbose up	1

# migrate down last one
migrate-down:
	migrate -path db/migration -database "postgresql://root:docker_postgres@localhost:5432/hrm_db?sslmode=disable" -verbose down 1

server:
	go run cmd/server/main.go	
