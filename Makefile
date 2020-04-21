
DIST_DIR=$(PWD)/dist
EXE_NAME=launchcode

all: clean build-linux build-mac build-windows

init: 
	mkdir -p ${DIST_DIR}

build-linux: init
	env GOOS=linux GOARCH=386 go build -o ${DIST_DIR}/${EXE_NAME}-linux-x86
	env GOOS=linux GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-linux-x64

build-mac: init
	env GOOS=darwin  GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-mac-x64

build-windows: init
	go generate
	env GOOS=windows GOARCH=386 go build -o ${DIST_DIR}/${EXE_NAME}-windows-x86
	env GOOS=windows GOARCH=amd64 go build -o ${DIST_DIR}/${EXE_NAME}-windows-x64

clean:
	if [ -d "${DIST_DIR}" ]; then rm -rvf ${DIST_DIR}; fi
	if [ -f resource.syso ]; then rm resource.syso; fi
