build_protos :
	protoc --proto_path /Users/user/Desktop/Test/protos --go_out=./internal/app/gen/ --go-grpc_out=./internal/app/gen/ /Users/user/Desktop/Test/protos/auth.proto
	protoc --proto_path /Users/user/Desktop/Test/protos --go_out=./internal/app/gen/ --go-grpc_out=./internal/app/gen/ /Users/user/Desktop/Test/protos/user.proto	
	protoc --proto_path /Users/user/Desktop/Test/protos --go_out=./internal/app/gen/ --go-grpc_out=./internal/app/gen/ /Users/user/Desktop/Test/protos/chat.proto

run_build: build_protos

run:
	./run_dev


