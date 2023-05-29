build-grpc:
	@echo " > Building [grpc]..."
	@cd ./cmd/grpc/ && go build && cd ../..
	@echo " > Finished building [grpc]"


run-grpc: build-grpc
	@echo " > Running [grpc]..."
	@cd ./cmd/grpc/ && ./grpc
	@echo " > Finished running [grpc]"

build-consumer:
	@echo " > Building [consumer]..."
	@cd ./cmd/mq/ && go build && cd ../..
	@echo " > Finished building [consumer]"