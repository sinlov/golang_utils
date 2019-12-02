#!/usr/bin/env bash

golang_x_pkg=(
 "glog"
"image"
"perf"
"snappy"
"term"
"sync"
"winstrap"
"cwg"
"leveldb"
"text"
"net"
"build"
"protobuf"
"dep"
"sys"
"crypto"
"gddo"
"tools"
"scratch"
"proposal"
"mock"
"oauth2"
"freetype"
"debug"
"mobile"
"gofrontend"
"lint"
"appengine"
"geo"
"review"
"arch"
"vgo"
"exp"
"time"
)

cd $GOPATH;

if [[ ! -d $GOPATH/src/golang.org/x ]]; then
	mkdir -p $GOPATH/src/golang.org/x
fi

ls;
# "define want you want install"
for name in "net" "build";do
   cd $GOPATH/src/golang.org/x;
   if [[ -d "$name" ]]
    then
     cd $name;
     echo $name "包已经存在,使用git pull来更新源码";
     git pull;
   else
     git_url="https://github.com/golang/${name}.git";
     echo "开始clone golang.org/x 在github.com上的镜像代码:${git_url}";
     git clone --depth 1 "$git_url";
     cd $name;
   fi
   #go install;
done
