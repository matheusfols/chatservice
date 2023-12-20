createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "postgres://postgres:root@host.docker.internal:5432/chat_test?sslmode=disable" -verbose up

migratedown:
	migrate -path=sql/migrations -database "postgres://postgres:root@host.docker.internal:5432/chat_test?sslmode=disable" -verbose drop	

grpc:
	protoc --go_out=. --go-grpc_out=. proto/chat.proto --experimental_allow_proto3_optional

.PHONY: migrate createmigration migratedown grpc