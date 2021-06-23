dir = $(shell pwd)
package = $(shell head -n 1 ./go.mod|cut -d ' ' -f 2 | rev | cut -d '/' -f 1 | rev)
version = 1.0.0

work = /data
go_path = /root/go
proto_path = /root/go/src
third_proto_path = /root/go/pkg/mod/github.com/go-kratos/kratos@v1.0.0/third_party
protoc_gen_bm = /root/go/bin/protoc-gen-bm
protoc_gen_go = /root/go/bin/protoc-gen-go
protoc_gen_gofast = /root/go/bin/protoc-gen-gofast

define rm
	if [ -n "$(shell docker ps -f "name=$(1)"| grep $(1) | cut -d ' ' -f 1)" ]; then docker stop $(1); fi
	if [ -n "$(shell docker ps -a -f "name=$(1)"| grep $(1) | cut -d ' ' -f 1)" ]; then docker rm $(1); fi
endef

.PHONY:build
build:
	docker build --network host -t debug:1.0.0 -f ./docker/go/Dockerfile .
	docker build --network host -t $(package):$(version) -f ./Debugfile .

.PHONY:init
init:
	$(call rm, $(package))
	make build
	docker run -it --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	bash ./init
	$(call rm, $(package))
	make build

.PHONY:bash
bash:
	$(call rm, $(package))
	docker run -it --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	/bin/bash

.PHONY:clean
clean:
	$(call rm, $(package))
	docker run -it --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	go clean --modcache

.PHONY:tidy
tidy:
	$(call rm, $(package))
	rm -rf ./go.sum
	docker run -it --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	go mod tidy

.PHONY:mod
mod: tidy
	$(call rm, $(package))
	rm -rf ./go.sum
	docker run -it --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	go mod vendor

.PHONY:install
install: build
	$(call rm, $(package))
	docker run --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	go build -o ./cmd/server ./cmd/main.go

.PHONY:run
run: install zipkin
	$(call rm, $(package))
	docker run --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	./cmd/server -conf ./configs/local -language ./language

.PHONY:proto
proto:
	$(call rm, $(package))
	docker run --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	bash ./protoc

.PHONY:wire
wire:
	$(call rm, $(package))
	docker run --name $(package) --net host \
	-v $(dir):$(work) \
	-v ~/go:$(go_path) \
	$(package):$(version) \
	bash ./wire

.PHONY:zipkin
zipkin:
	$(call rm, "zipkin")
	docker run -d --name zipkin --net host openzipkin/zipkin:2

.PHONY:jaeger
jaeger:
	$(call rm, "jaeger")
	docker run --name jaeger \
	--net host \
	jaegertracing/all-in-one:1.22.0
