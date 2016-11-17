#!/usr/bin/env bash

buildPath="build"
ReNameCode="golang_utils"
VERSION_MAJOR=0
VERSION_MINOR=0
VERSION_PATCH=1
VERSION_BUILD=0

VersionCode=$[$[VERSION_MAJOR * 100000000] + $[VERSION_MINOR * 100000] + $[VERSION_PATCH * 100] + $[VERSION_BUILD]]
VersionName="${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}.${VERSION_BUILD}"

shell_running_path=$(cd `dirname $0`; pwd)

if [ -d "${buildPath}" ]; then
    rm -rf ${buildPath}
    sleep 1
fi

mkdir -p ${buildPath}

echo -e "============\nPrint build info start"
go version
which go
echo -e "Your settings is
\tVersion Name -> ${ReNameCode}
\tVersion code -> ${VersionCode}
\tVersion name -> ${VersionName}
\tOut Path -> ${shell_running_path}/${buildPath}
"
echo -e "Print build info end\n============"

echo "start build OSX 64"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
mv main "${buildPath}/${ReNameCode}_${VersionName}_${VersionCode}_osx_64"
echo "build OSX 64 finish"

echo "start build Linux 64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
mv main "${buildPath}/${ReNameCode}_${VersionName}_${VersionCode}__linux_64"
echo "build linux 64 finish"

echo "start build windows 64"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
mv main.exe "${buildPath}/${ReNameCode}__${VersionName}_${VersionCode}_win_86_64.exe"
echo "build windows 64 finish"

echo -e "============\nall the build is finish! at Path\n${shell_running_path}\\${buildPath}"