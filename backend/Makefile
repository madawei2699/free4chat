GO ?= go
GOPATH := $(CURDIR)/src/_vendor:$(GOPATH)

all: clean build

clean:
	rm -rf build/

build:
	mkdir build && cd ./src/ && $(GO) build -o ../build/app