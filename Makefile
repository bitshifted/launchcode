
DIST_DIR=$(PWD)/dist
EXE_NAME=launchcode
GOPATH := $(shell go env GOPATH)
PROJECT_ROOT=${PWD}
TARGET_DIR := ${PROJECT_ROOT}/target

all: clean build-linux build-mac build-windows

init-launchcode: 
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	mkdir -p ${DIST_DIR}

init-launch-config:
	mkdir ${TARGET_DIR}

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GOPATH)/bin/golangci-lint run ./...
	go mod tidy

test: init-launch-config
	go test  -coverprofile=${TARGET_DIR}/coverage.out ./...
	go tool cover -html=${TARGET_DIR}/coverage.out -o ${TARGET_DIR}/coverage.html

build-launch-config: clean init-launch-config test
	cd cmd/launch-config && \
	env GOOS=linux GOARCH=amd64 go build -trimpath -buildvcs=false -o ${TARGET_DIR}/launch-config

build-linux: init-launchcode
	env GOOS=linux GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-linux-x64
	env GOOS=linux GOARCH=arm64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-linux-aarch64

build-mac: init-launchcode
	env GOOS=darwin  GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-mac-x64
	env GOOS=darwin  GOARCH=arm64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-mac-aarch64

build-windows: init-launchcode
	go generate ./...
	env GOOS=windows GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-windows-x64.exe


clean:
	if [ -d "${DIST_DIR}" ]; then rm -rvf ${DIST_DIR}; fi
	if [ -d "${TARGET_DIR}" ]; then rm -rvf ${TARGET_DIR}; fi
	if [ -f resource.syso ]; then rm resource.syso; fi
