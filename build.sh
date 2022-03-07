#!/usr/bin/env bash

basedir=`cd $(dirname $0); pwd -P`

# 项目名称
PROJECT=go-fragments
# 版本号
VERSION=v0.1
# 编译结果路径
RELEASES=releases

RELEASES_DIR=${basedir}/${RELEASES}/${VERSION}

function init() {
  rm -fr "${RELEASES_DIR}"
  echo "clean original releases."
}

function prepare() {
  mkdir -p "${RELEASES_DIR}"
  echo "make releases dir."
}

function build() {
    os=$1
    arch=$2
    name=$3
    env=$4

    GOOS=${os} GOARCH=${arch} go build -o "${RELEASES_DIR}/${name}" -ldflags "-X main.ENV=${env} -X main.VERSION=${VERSION}"
    echo "build ${name} finish."
}

cd ${basedir}/src

build windows amd64  ${PROJECT}-win64-dev-${VERSION}.exe dev
build windows amd64 ${PROJECT}-win64-${VERSION}.exe online
build darwin amd64 ${PROJECT}-darwin-dev-${VERSION} dev
build darwin amd64 ${PROJECT}-darwin-${VERSION} online

cd ${basedir}