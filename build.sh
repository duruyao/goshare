#!/usr/bin/env bash

## date:   2021.11.16
## author: duruyao@gmail.com
## desc:   cross compile GoShare for multi-platform
## usage:  bash build.sh [all]

set -euo pipefail

release_list=(
  'GoShare-macOS-amd64 darwin amd64 goshare'
  'GoShare-Linux-386 linux 386 goshare'
  'GoShare-Linux-arm linux arm goshare'
  'GoShare-Linux-amd64 linux amd64 goshare'
  'GoShare-Windows-386 windows 386 goshare.exe'
  'GoShare-Windows-arm windows arm goshare.exe'
  'GoShare-Windows-amd64 windows amd64 goshare.exe'
)

package="github.com/duruyao/goshare"

## compile goshare for current platform
if [ -z "$1" ] || [ "$1" != "all" ]; then
  bash -x -c "GOROOT=${GOROOT} GOPATH=${GOPATH} GO_ENABLED=0 ${GOROOT}/bin/go build -o ${PWD}/goshare ${package}"
  exit
fi

## cross compile goshare for multi-platform
releases_dir="${PWD}/releases"
version="$(grep -o "version.*" "${PWD}"/package.json | grep -o "[0-9]\+.[0-9]\+.[0-9]\+")"

mkdir -p "${releases_dir}"
rm -rf "${releases_dir:?}"/*
pushd "${releases_dir}" 1>/dev/null 2>&1

for release in "${release_list[@]}"; do
  # shellcheck disable=SC2206
  release=(${release})
  deploy_dir="${PWD}/${release[0]}-${version}"
  mkdir -p "${deploy_dir}"

  bash -x -c "GOROOT=${GOROOT} GOPATH=${GOPATH} GO_ENABLED=0 GOOS=${release[1]} GOARCH=${release[2]} ${GOROOT}/bin/go build -o ${deploy_dir}/${release[3]} ${package}"
  chmod +x "${deploy_dir}/${release[3]}"

  if [ -n "$(command -v zip)" ]; then
    bash -x -c "zip -r ${deploy_dir}.zip $(basename "${deploy_dir}") 1>/dev/null"
  fi

  if [ -n "$(command -v tar)" ]; then
    bash -x -c "tar -cvf ${deploy_dir}.tar $(basename "${deploy_dir}") 1>/dev/null"
    bash -x -c "tar -zcvf ${deploy_dir}.tar.gz $(basename "${deploy_dir}") 1>/dev/null"
    bash -x -c "tar -jcvf ${deploy_dir}.tar.bz2 $(basename "${deploy_dir}") 1>/dev/null"
  fi

  echo ""
done

popd 1>/dev/null 2>&1
