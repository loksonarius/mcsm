#!/usr/bin/env bash
set -e

build_dir="build"
build_local="${1:-no}"
package_name=${BINARY}
version_flags="-X main.version=${VERSION} -X main.commit=${COMMIT}"
linker_flags='-s -w -extldflags "-static"'

mkdir -p "${build_dir}"

case "${build_local}" in

  "yes"|"true"|"YES"|"yes please")
    go build \
      -o "${build_dir}/${package_name}-local" \
      -i -trimpath \
      -ldflags "${version_flags} ${linker_flags}"
  ;;

  *)
    archs=("amd64" "arm" "arm64" "mips" "mips64")
    for GOARCH in "${archs[@]}"
    do
      output_name="${build_dir}/${package_name}-linux-${GOARCH}"

      env GOOS="linux" GOARCH="${GOARCH}" \
        go build \
          -o "${output_name}" \
          -i -trimpath \
          -ldflags "${version_flags} ${linker_flags}"
    done
  ;;
esac

echo 'Built'
