.PHONY: gen-pb
gen-pb: ## Generate protobuf files
	@echo Starting generation of protobuf files
	@protoc --proto_path=./proto \
             --go_out=paths=source_relative:./proto \
             --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./proto \
             ./proto/*.proto
	@echo Successfully generated protobuf files

	@echo Starting to inject tags
	@protoc-go-inject-tag -input="./proto/*.pb.go"
	@echo Successfully injected tags

	@echo Starting to generate mocks
	@mockery --dir=./proto --all --output=./mocks
	@echo Successfully generated mocks
