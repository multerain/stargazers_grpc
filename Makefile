REPO_ROOT_DIR = $(shell pwd)

.PHONY: protobuf
protobuf:
	protoc -I ${REPO_ROOT_DIR}/stargazers --proto_path=${GOPATH} --proto_path=${REPO_ROOT_DIR} \
		--go_out=${REPO_ROOT_DIR}/stargazers --go_opt=paths=source_relative --go-grpc_out=${REPO_ROOT_DIR}/stargazers \
		--go-grpc_opt=paths=source_relative stargazers/stargazers.proto