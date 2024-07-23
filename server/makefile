PROTO_PATH=/Users/user/Desktop/Test/protos

build_ent:
	go generate ./internal/auth/gen/ent
	go generate ./internal/user/gen/ent
	go generate ./internal/chat/gen/ent

build_protos :
	protoc --proto_path $(PROTO_PATH) --go_out=./pkg/gen/ --go-grpc_out=./pkg/gen/ $(PROTO_PATH)/auth.proto
	protoc --proto_path $(PROTO_PATH) --go_out=./internal/auth/gen/ --go-grpc_out=./internal/auth/gen/ $(PROTO_PATH)/user.proto
	protoc --proto_path $(PROTO_PATH) --go_out=./internal/user/gen/ --go-grpc_out=./internal/user/gen/ $(PROTO_PATH)/user.proto	
	protoc --proto_path $(PROTO_PATH) --go_out=./internal/chat/gen/ --go-grpc_out=./internal/chat/gen/ $(PROTO_PATH)/chat.proto

run_build: build_ent build_protos

run:
	docker compose up --build --remove-orphans

create_migration :
	atlas migrate diff $(name) \
		--dir "file://migrations/$(feature)" \
		--to "ent://internal/$(feature)/gen/ent/schema" \
		--dev-url "docker://postgres/16/$(feature)?search_path=public"


