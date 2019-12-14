.DEFAULT_GOAL := dev

GO111MODULE := on

CMD := bofh

BIN_DIR := ./bin
INSTALL_DIR ?= /usr/local/bin

init:
	go mod init

build:
	mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${CMD} ./...

verify:
	go mod verify

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

dev: build verify tidy fmt vet test

install: dev
	go install ./...

install-bin: dev
	cp ${BIN_DIR}/${CMD} ${INSTALL_DIR}

uninstall-bin:
	rm -f ${INSTALL_DIR}/${CMD}

clean:
	rm -rf ${BIN_DIR}/*

.PHONY: init build verify tidy fmt vet test
	    dev install install-bin uninstall-bin clean
