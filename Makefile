all: generate

.PHONY: proto

generate:
	echo "proto->golang"
	docker run --rm -v ${PWD}/proto/definitions:/defs -v ${PWD}/proto/generated:/gen namely/protoc-all:1.47_2 -d /defs -l go -o /gen --with-validator