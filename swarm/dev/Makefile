.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/rupayaproject/go-rupaya/cmd/swarm
	go build -o bin/rupaya     github.com/rupayaproject/go-rupaya/cmd/rupaya
	go build -o bin/bootnode github.com/rupayaproject/go-rupaya/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/rupayaproject/go-rupaya/swarm/...
