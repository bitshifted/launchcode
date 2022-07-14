
DIST_DIR=$(PWD)/dist
EXE_NAME=launchcode

all: clean build-linux build-mac build-windows

init: 
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	mkdir -p ${DIST_DIR}

build-linux: init
	env GOOS=linux GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-linux

build-mac: init
	env GOOS=darwin  GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-mac

build-windows: init
	go generate ./...
	env GOOS=windows GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-windows.exe

clean:
	if [ -d "${DIST_DIR}" ]; then rm -rvf ${DIST_DIR}; fi
	if [ -f resource.syso ]; then rm resource.syso; fi
