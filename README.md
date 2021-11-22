# bugsnag-exporter

![Go Workflow](https://github.com/xander-jones/bugsnag-exporter/actions/workflows/go.yml/badge.svg)
![Latest Stable Version](https://img.shields.io/github/v/release/xander-jones/bugsnag-exporter)

Tool for extracting data from Bugsnag's Data Access API and turning it into a JSON or CSV file.

:warning: **Work In Progress.**

## Usage

### Download and install

First, download a copy of the binary from the [latest release](https://github.com/xander-jones/bugsnag-exporter/releases). There are different versions of this binary depending on your platform. All binaries are built for 64-bit:


| Platform | Binary variant |
|---|---|
| macOS Intel | `bugsnag-exporter-<VERSION>-darwin-amd64` |
| macOS Apple Silicon | `bugsnag-exporter-<VERSION>-darwin-arm64` |
| Linux | `bugsnag-exporter-<VERSION>-linux-amd64` |
| Windows | `bugsnag-exporter-<VERSION>-windows-amd64.exe` |

For Linux and macOS, you will need to set this binary to be executable with

```sh
$ chmod +x BINARY_FILENAME
```

### Generate Bugsnag API token

From your Bugsnag dashboard, [generate a new Personal Authentication Token](https://app.bugsnag.com/settings/my-account). *Note*, this is not the same as your project API key.

Save this token, you will need this to authenticate any and all calls to the Data Access API with the `--token TOKEN` parameter.


### Help file

To see the `bugsnag-exporter` manual pages, run `$ ./bugsnag-exporter --help`

```sh
Usage of /var/folders/j3/gpggwqd53wj_z5mhtn9yf2t40000gn/T/go-build962128733/b001/exe/main:
  -csv
        [Flag] Output data to file as CSV. Default false, noramally outputs as JSON
  -error-id string
        [String] An error ID to download. If provided, downloads all events within filters for this error ID
  -events
        [Flag] Download events rather than error groups when this flag is enabled. Requires --project-id (and optionally --error-id)
  -filters string
        [String] A string array of filters to apply (URL format)
  -minimal
        [Flag] Download minimal event reports only for smaller file sizes
  -no-warn
        [Flag] Don't warn me if this call will take more than 5 calls to the API
  -output-dir string
        [String] Directory to store the downloaded file (default "./data")
  -project-id string
        [String] The Project ID you wish to download from
  -show-project-ids
        [Flag] Use this flag to get a list of project IDs accessible with your token. Overrides any other flags
  -token string
        [REQUIRED][String] Your Bugsnag personal auth token
  -verbose
        [Flag] Set the output to be verbose for debugging purposes.
```

### Getting data from the API



#### Get a list of projects accessible for your API token

```
$ ./bugsnag-exporter --token TOKEN 
```

#### JSON or CSV

By default, data will be saved in a JSON format. This is most appropriate for the type of data downloaded from Bugsnag as it has multi-level keys. However, if you wish to save data as a comma seperated values file (`*.csv`), you can add the `--csv` flag to do so.

CSV files will have a single row per array elements, and single columns per top-level key. Sub-keys for an element will be stored in plain-text within the cell, for example, the following is a JSON <-> CSV equivalent:

##### JSON

```json
[
  {
    "keyA": {
      "keyA1": "el1valueA1",
      "keyA2": "el1valueA2"
    },
    "keyB": {
      "keyB1": "el1valueB1",
      "keyB2": "el1valueB2"
    }
  },
  {
    "keyA": {
      "keyA1": "el2valueA1",
      "keyA2": "el2valueA2"
    },
    "keyB": {
      "keyB1": "el2valueB1",
      "keyB2": "el2valueB2"
    }
  }
]
```

##### CSV

```csv
keyA,keyB
"{\"keyA1\":\"el1valueA1\",\"keyA2\":\"el1valueA2\"}","{\"keyA1\":\"el2valueA1\",\"keyA2\":\"el2valueA2\"}",
```

or visualized as a table:

|KeyA|KeyB|
|---|---|
|{\"keyA1\":\"el1valueA1\",\"keyA2\":\"el1valueA2\"}|{\"keyA1\":\"el2valueA1\",\"keyA2\":\"el2valueA2\"}|
#### Filtering



#### Extra options

* `--no-warn`: don't show a warning if more than 5 API calls are going to be made for this command.
* 

## Building from source

#### Build for all main platforms and architectures

Run the `make.sh` script. This will generate builds in a `./build` directory for:

```sh
os_archs=(
    darwin/amd64   # macOS Intel 64-bit
    darwin/arm64   # macOS Apple Silicon 64-bit
    linux/amd64    # Linux 64-bit
    windows/amd64  # Windows 64-bit
)
```

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
