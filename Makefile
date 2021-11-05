generate:
	rm -rf gen/{pb,clean_proto,swagger}/{*.go,*.proto,*.json}
	mkdir -p gen/{pb,clean_proto,swagger}

	protoc \
		-I=proto/ \
		-I=${GOPATH}/src/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=./ \
		--gorm_out=./ \
		--validate_out=lang=go:./ \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=allow_merge=true,merge_file_name=api:./gen/swagger/ proto/*.proto

	protoc \
		-I=proto/ \
		-I=${GOPATH}/src/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go-grpc_out=./ proto/api.proto proto/marks.proto proto/exports.proto

	go run scripts/generate_clean_protos/main.go

	gofmt -s -w gen/pb/

	make lint --ignore-errors > /dev/null 2>&1

lint:
	protolint lint -config_path=protolint.yaml -fix gen/clean_proto
