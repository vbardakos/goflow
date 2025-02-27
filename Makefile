proto:
	protoc --go_out=. --go-drpc_out=. --go_opt=paths=source_relative engine/engine.proto

.PHONY: proto
