#! /bin/bash
#
# /*
#  * Copyright (c) 2023  Bitshift D.O.O (http://bitshifted.co)
#  *
#  * This Source Code Form is subject to the terms of the Mozilla Public
#  * License, v. 2.0. If a copy of the MPL was not distributed with this
#  * file, You can obtain one at http://mozilla.org/MPL/2.0/.
#  */
#

PATCH_REGEX='^(build|chore|ci|docs|fix|perf|refactor|revert|style|test)\s?(\(.+\))?\s?:\s*(.+)'
MINOR_REGEX='^(feat)\s*(\(.+\))?\s?:\s*(.+)'
MAJOR_REGEX='^(BREAKING CHANGE)\s*(\(.+\))?\s?:\s*(.+)'

# get the latest tag
LATEST_TAG=$(git describe --tags `git rev-list --tags --max-count=1`  2> /dev/null)
if [ -z $LATEST_TAG ]; then
  LATEST_TAG="v1.0.0"
  echo "1.0.0"
  exit 0
fi

LATEST_VERSION=""
if [[ $LATEST_TAG =~ [0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    LATEST_VERSION=${BASH_REMATCH[0]}
else
    echo "Failed to extract current version" >&2
    exit 1
fi

SCRIPT_DIR=$( dirname -- "$0"; )
# process last commit message
MESSAGE=$(git log -1 --pretty=%B)
if [[ "$MESSAGE" =~ $PATCH_REGEX ]]; then
  "$SCRIPT_DIR"/semver-bump.sh -p $LATEST_VERSION
elif [[ "$MESSAGE" =~ $MINOR_REGEX ]]; then
  "$SCRIPT_DIR"/semver-bump.sh -m $LATEST_VERSION
elif [[ $MESSAGE =~ $MAJOR_REGEX ]]; then
  "$SCRIPT_DIR"/semver-bump.sh -M $LATEST_VERSION
else
  "$SCRIPT_DIR"/semver-bump.sh -p $LATEST_VERSION
fi




