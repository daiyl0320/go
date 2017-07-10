#!/bin/sh

# Save the pwd before we run anything
PRE_PWD=`pwd`

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Setup the environment for the build
export GOPATH=$GOPATH:$BUILD_DIR

# Build the project
cd $BUILD_DIR
go clean
go build -o ../bin/entry entry
go install entry

EXIT_STATUS=$?

if [ $EXIT_STATUS == 0 ]; then
  echo "Build succeeded"
else
  echo "Build failed"
fi

# Change back to where we were
cd $PRE_PWD

exit $EXIT_STATUS
