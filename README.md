# bugsnag-to-csv

![Go Workflow](https://github.com/xander-jones/bugsnag-to-csv/actions/workflows/go.yml/badge.svg)

Tool for extracting data from Bugsnag's Data Access API and turning it into a CSV.

:warning: **Work In Progress.**

## Usage

```
./build/darwin/amd64/bugsnag-to-csv --help
Usage of ./build/darwin/amd64/bugsnag-to-csv:
  -show-project-ids
        Use this flag to get a list of project IDs accessible with your token.
  -token string
        [REQUIRED][String] Your Bugsnag personal auth token.
```

## Building

#### Build for all main platforms and architectures

Run the `make.sh` script. This will generate builds in a `./build` directory:

```sh
$ ./make.sh
```

#### Building for a specific platform and architecture

Replace `YOUR_OS` and `YOUR_ARCH` in the following command and run to build an executable binary for a custom setup:

```sh
$ env GOOS=YOUR_OS GOARCH=YOUR_ARCH go build .
```

To see which OS & architectures your Go installation can build for you can run:

```sh
$ go tool dist list
```
