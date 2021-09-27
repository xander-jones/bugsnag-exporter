# bugsnag-exporter

![Go Workflow](https://github.com/xander-jones/bugsnag-exporter/actions/workflows/go.yml/badge.svg)
![Latest Stable Version](https://img.shields.io/github/v/release/xander-jones/bugsnag-exporter)

Tool for extracting data from Bugsnag's Data Access API and turning it into a CSV.

:warning: **Work In Progress.**

## Usage

### Using source

```
$ go run main.go [params]
```

### Using the binary

```
$ ./build/darwin/amd64/bugsnag-exporter --help
Usage of ./build/darwin/amd64/bugsnag-exporter:
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
