#!/bin/bash
#
# Builds binaries for all of the architectures and platforms listed
# in `os_archs`, call `./make.sh` to run this bash script.
#
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
#

os_archs=(
    darwin/amd64   # macOS Intel 64-bit
    darwin/arm64   # macOS Apple Silicon 64-bit
    linux/amd64    # Linux 64-bit
    windows/amd64  # Windows 64-bit
)

build_success=()
build_failures=()

printf "Building bugsnag-exporter binaries in ./build for the chosen platforms & architectures\r\n\r\n"
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
    outputfile=./build/$os_arch/bugsnag-exporter
    if [ $goos == "windows" ]; then
        outputfile="$outputfile.exe"
    fi
    GOOS=${goos} GOARCH=${goarch} go build -o $outputfile main.go &> /dev/null
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