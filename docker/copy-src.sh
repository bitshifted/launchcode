!# /bin/bash

if [ -z $1 ];then
    echo "Usage: copy-src <targte-directory>"
    exit 1
fi

TARGET_DIR=$1
cp -rv /usr/src/launchcode $TARGET_DIR
echo "Source code copied successfully"