proto:
	protoc --go_out=. --go-drpc_out=. --go_opt=paths=source_relative actor/actor.proto

tests:
	@ginkgo -r

.PHONY: proto
