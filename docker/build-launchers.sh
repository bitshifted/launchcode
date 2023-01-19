!# /bin/bash

if [ -z $1 ];then
    echo "Usage: build-launchers <targte-directory>"
    exit 1
fi

TARGET_DIR=$1
cd $TARGET_DIR
make all
echo "launchers built successfully"