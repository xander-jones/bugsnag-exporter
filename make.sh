#!/bin/bash

# https://github.com/golang/go/blob/master/src/go/build/syslist.go

os_archs=(
    darwin/386
    darwin/amd64
    linux/386
    linux/amd64
    linux/arm
    linux/arm64
    linux/ppc64
    linux/ppc64le
    linux/mips
    linux/mipsle
    linux/mips64
    linux/mips64le
    linux/riscv64
    linux/s390x
    windows/386
    windows/amd64
    windows/arm
)

build_success=()
build_failures=()

printf "Building bugsnag-to-csv binaries for the chosen platforms & architectures in ./build\r\n\r\n"
# clear out the ./build directory to begin with
# and then recreate it so we have a fresh output.
if [ -d "build" ]; then
    rm -rf build
fi
mkdir "build"

idx=1
for os_arch in "${os_archs[@]}"
do
    printf "\33[2K\rBuilding binary %s of ${#os_archs[@]} ($os_arch)" $idx
    goos=${os_arch%/*}
    goarch=${os_arch#*/}
    GOOS=${goos} GOARCH=${goarch} go build -o "./build/$os_arch/bugsnag-to-csv" main.go &> /dev/null
    if [ $? -eq 0 ]
    then
        build_success+=(${os_arch})
    else
        build_failures+=(${os_arch})
    fi
    ((idx=idx+1))
done

printf "\33[2K\rSuccesfuly built for:\r\n"
for os_arch in "${build_success[@]}"
do
    printf "\t%s\n" "${os_arch}"
done
echo

printf "Failed to build for:\r\n"
for os_arch in "${build_failures[@]}"
do
    printf "\t%s\n" "${os_arch}"
done
echo