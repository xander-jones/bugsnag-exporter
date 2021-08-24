## Building for release

### Build for all main platforms and architectures

```sh
$ ./make.sh
```

### Building for macOS

```sh
$ env GOOS=darwin GOARCH=386 go build .
```

Or, use a different `GOARCH` for your system's architecture.

### Building foe Windows

```sh
$ env GOOS=windows GOARCH=386 go build .
```

### Building for Linux

