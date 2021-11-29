#!/usr/bin/env bash

# date:   2021.11.16
# author: duruyao@gmail.com
# desc:   cross compile GoFS for multi-platform

releases=(
  'GoFS-macOS-amd64 darwin amd64'
  'GoFS-Linux-386 linux 386'
  'GoFS-Linux-arm linux arm'
  'GoFS-Linux-amd64 linux amd64'
  'GoFS-Windows-386.exe windows 386'
  'GoFS-Windows-arm.exe windows arm'
  'GoFS-Windows-amd64.exe windows amd64')

echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
GOEXEC="$(command -v go)"
echo "GOEXEC=${GOEXEC}"
echo

releases_dir="$(pwd)/releases"

mkdir -p "${releases_dir}"

for release in "${releases[@]}"; do
  items=(${release}) # NOTE: (${release}) != ("${release}")
  #  echo "${items[*]}"
  #  echo "0: ${items[0]}"
  #  echo "1: ${items[1]}"
  #  echo "2: ${items[2]}"
  cmd="GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o ${releases_dir}/${items[0]} github.com/duruyao/gofs"
  echo "${cmd}"
  GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o "${releases_dir}"/"${items[0]}" github.com/duruyao/gofs
done
