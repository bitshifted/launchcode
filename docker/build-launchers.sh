#! /bin/bash

WORKDIR=${PWD}
LAUNCHCODE_SRC_DIR=/usr/src/launchcode
LAUNCHCODE_EMBED_DIR=$LAUNCHCODE_SRC_DIR/config/embed
LINUX_CONFIG_DIR=${PWD}/output/linux
MAC_CONFIG_DIR=${PWD}/output/mac
WINDOWS_CONFIG_DIR=${PWD}/output/windows
DIST_DIR=dist
DEFAULT_ICON="launchcode.ico"

if [ -z $1 ];then
    echo "Usage: build-launchers <config-file>"
    exit 1
fi

echo "Using configuration file $1"
/usr/bin/launch-config $1
echo "Configuration created successfully!"

cd $LAUNCHCODE_SRC_DIR
echo "Initializing launchers build environment"
make init-launchcode
echo "Building Linux launchers..."
rm -v $LAUNCHCODE_EMBED_DIR/*
cp $LINUX_CONFIG_DIR/* $LAUNCHCODE_EMBED_DIR
make build-linux
echo "Linux launcher built successfully!"

echo "Building Mac OS launchers..."
rm -v $LAUNCHCODE_EMBED_DIR/*
cp $MAC_CONFIG_DIR/* $LAUNCHCODE_EMBED_DIR
make build-mac
echo "Mac OS launcher built sucessfully"

echo "Building Windows launchers..."
rm -v $LAUNCHCODE_EMBED_DIR/*
cp $WINDOWS_CONFIG_DIR/* $LAUNCHCODE_EMBED_DIR
cp $LAUNCHCODE_SRC_DIR/versioninfo.json $LAUNCHCODE_SRC_DIR/cmd/launchcode
# setup icons
ICON_FILE=$(yq '.windows.icon' $1)
if [ -z $ICON_FILE ];then
    echo "Using default Windows icon"
    ICON_FILE=$DEFAULT_ICON
    cp icons/$DEFAULT_ICON cmd/launchcode
else
    echo "Windows icon file: $ICON_FILE"
    cp $ICON_FILE $LAUNCHCODE_SRC_DIR/cmd/launchcode
fi
jq  --arg icon_file $ICON_FILE '.IconPath |= $icon_file' versioninfo.json > $LAUNCHCODE_SRC_DIR/cmd/launchcode/versioninfo.json
make build-windows
echo "Windows launcher built sucessfully"

cd $WORKDIR
cp -rv /usr/src/launchcode/$DIST_DIR .
OWNER_ID=$(stat -c %u $1)
GROUP_ID=$(stat -c %g $1)
chown -R $OWNER_ID:$GROUP_ID $DIST_DIR

# cleanup
rm -rvf output