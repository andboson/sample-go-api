export GOPATH=$(CURDIR)/../../
NOVENDOR=`glide novendor`

all: godeps test  build

godeps:
	rm -rf ${GOPATH}/pkg/*
	glide install

test:
	echo test dirs: ${NOVENDOR}
	go test -v ${NOVENDOR}

build:
	echo GOPATH is ${GOPATH}
	go build -v