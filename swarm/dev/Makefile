.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/rupayaproject/rupaya/cmd/swarm
	go build -o bin/rupaya     github.com/rupayaproject/rupaya/cmd/rupaya
	go build -o bin/bootnode github.com/rupayaproject/rupaya/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/rupayaproject/rupaya/swarm/...
