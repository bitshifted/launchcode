!# /bin/bash

if [ -z $1 ];then
    echo "Usage: build-launchers <target-directory>"
    exit 1
fi

TARGET_DIR=$1
cd $TARGET_DIR
make all
chown -Rv 1000:1000 $TARGET_DIR/*
echo "launchers built successfully"