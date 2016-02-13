export GOPATH=$(CURDIR)/../../

all: godeps test  build

godeps:
	rm -rf ${GOPATH}/pkg/*
    glide install

test:
	go test -v $(glide novendor)

build:
	echo GOPATH is ${GOPATH}
	go build -v