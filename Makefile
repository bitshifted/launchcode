
DIST_DIR=$(PWD)/dist
EXE_NAME=launchcode

all: clean build-linux build-mac build-windows

init: 
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	mkdir -p ${DIST_DIR}

build-linux: init
	env GOOS=linux GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-linux-x64
	env GOOS=linux GOARCH=arm64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-linux-aarch64

build-mac: init
	env GOOS=darwin  GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-mac-x64
	env GOOS=darwin  GOARCH=arm64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-mac-aarch64

build-windows: init
	go generate ./...
	env GOOS=windows GOARCH=amd64 go build -trimpath  -buildvcs=false -o ${DIST_DIR}/${EXE_NAME}-windows-x64.exe

test:
	go test ./...

clean:
	if [ -d "${DIST_DIR}" ]; then rm -rvf ${DIST_DIR}; fi
	if [ -f resource.syso ]; then rm resource.syso; fi
