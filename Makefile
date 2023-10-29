OUT := build/imgsrt
PKG := github.com/RATIU5/imgsrt
VERSION := 0.0.1
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: run

server:
	go build -v -o ${OUT} -ldflags="-X github.com/RATIU5/imgsrt/internal.Version=${VERSION}" ${PKG}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	go build -i -v -o ${OUT}-v${VERSION} -ldflags="-extldflags \"-static\" -w -s -X github.com/RATIU5/imgsrt/internal.Version=${VERSION}" ${PKG}

run: server
	./${OUT}

clean:
	-@rm ${OUT} ${OUT}-v*
