#!/usr/bin/env bash

## date:   2021.11.16
## author: duruyao@gmail.com
## desc:   cross compile GoFS for multi-platform

releases=(
  'GoFS-macOS-amd64 darwin amd64 GoFS-macOS-amd64 gofs'
  'GoFS-Linux-386 linux 386 GoFS-Linux-386 gofs'
  'GoFS-Linux-arm linux arm GoFS-Linux-arm gofs'
  'GoFS-Linux-amd64 linux amd64 GoFS-Linux-amd64 gofs'
  'GoFS-Windows-386.exe windows 386 GoFS-Windows-386 gofs.exe'
  'GoFS-Windows-arm.exe windows arm GoFS-Windows-arm gofs.exe'
  'GoFS-Windows-amd64.exe windows amd64 GoFS-Windows-amd64 gofs.exe'
)

echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
GOEXEC="$(command -v go)"
echo "GOEXEC=${GOEXEC}"
echo

releases_dir="$(pwd)/releases"
version_id="$(date '+%Y.%m.%d')"
if [ -n "$1" ]; then
  version_id="$1"
fi

mkdir -p "${releases_dir}" && rm -rf "${releases_dir}"/GoFS*

pushd "${releases_dir}" || exit

for release in "${releases[@]}"; do
  items=(${release}) ## NOTE: ( ${release} ) != ( "${release}" )
  cmd="GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o ${releases_dir}/${items[0]} github.com/duruyao/gofs"
  echo "${cmd}"
  GO_ENABLED=0 GOOS=${items[1]} GOARCH=${items[2]} ${GOEXEC} build -o "${releases_dir}"/"${items[0]}" github.com/duruyao/gofs
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
