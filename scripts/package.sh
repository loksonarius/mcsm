#!/usr/bin/env bash
set -e

build_dir="build"
package_name=${BINARY}

archs=("amd64" "arm" "arm64" "mips" "mips64")
for arch in "${archs[@]}"; do
  binary_path="${build_dir}/${package_name}-linux-${arch}"

  if [[ -f "${binary_path}" ]]; then
    tar_path="${binary_path}.tgz"
    cp "${binary_path}" mcsm
    tar -zcf "${tar_path}" mcsm
    rm mcsm
  fi
done

echo 'Packaged'
