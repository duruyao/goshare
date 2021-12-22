#!/usr/bin/env bash

## date:   2021.11.16
## author: duruyao@gmail.com
## desc:   cross compile GoShare for multi-platform

releases=(
  'goshare-macos-amd64 darwin amd64 GoShare-macOS-amd64 goshare'
  'goshare-linux-386 linux 386 GoShare-Linux-386 goshare'
  'goshare-linux-arm linux arm GoShare-Linux-arm goshare'
  'goshare-linux-amd64 linux amd64 GoShare-Linux-amd64 goshare'
  'goshare-windows-386.exe windows 386 GoShare-Windows-386 goshare.exe'
  'goshare-windows-arm.exe windows arm GoShare-Windows-arm goshare.exe'
  'goshare-windows-amd64.exe windows amd64 GoShare-Windows-amd64 goshare.exe'
)

echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
GOEXEC="$(command -v go)"
echo "GOEXEC=${GOEXEC}"
echo

## compile goshare for current platform
if [ -z "$1" ] || [ "$1" != "all" ]; then
  echo "GO_ENABLED=0 ${GOEXEC} build github.com/duruyao/goshare"
  GO_ENABLED=0 ${GOEXEC} build github.com/duruyao/goshare
  exit
fi

## cross compile goshare for multi-platform
releases_dir="${PWD}/releases"
version_id="$(date '+%Y.%m.%d')"
if [ -n "$2" ]; then
  version_id="$2"
fi

mkdir -p "${releases_dir}" && rm -rf "${releases_dir:?}"/*

pushd "${releases_dir}" || exit

for release in "${releases[@]}"; do
  items=(${release}) ## NOTE: ( ${release} ) != ( "${release}" )
  cmd="GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o ${releases_dir}/${items[0]} github.com/duruyao/goshare"
  echo "${cmd}"
  GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o "${releases_dir}"/"${items[0]}" github.com/duruyao/goshare
  deploy_dir="${releases_dir}/${items[3]}-${version_id}-release"

  mkdir -p "${deploy_dir}"
  cp "${releases_dir}/${items[0]}" "${deploy_dir}/${items[4]}"

  if [ -n "$(command -v zip)" ]; then
    zip -r "${deploy_dir}.zip" "$(basename "${deploy_dir}")"
  fi

  if [ -n "$(command -v tar)" ]; then
    tar -cvf "${deploy_dir}.tar" "$(basename "${deploy_dir}")"
    tar -zcvf "${deploy_dir}.tar.gz" "$(basename "${deploy_dir}")"
    tar -jcvf "${deploy_dir}.tar.bz2" "$(basename "${deploy_dir}")"
  fi

  rm -r "${deploy_dir}" || exit
done

popd || exit
