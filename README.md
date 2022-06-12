# iogo 

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/Zarthus/iogo.svg)](https://pkg.go.dev/github.com/Zarthus/iogo/v2)
[![Coverage Status](https://coveralls.io/repos/github/Zarthus/iogo/badge.svg?branch=main)](https://coveralls.io/github/Zarthus/iogo?branch=main)

`iogo` is a helper go library for handling input and output, and stylizing them.

At it's core, iogo is a reader and a writer that doesn't do a whole more than
`fmt.Print` and `fmt.Scanf` with pretty styling options, and if you don't need that,
you probably don't need this library.

Where iogo helps you, is by offering an extensive toolkit for
terminal colours, handling input, and providing styling for output.

[![cast](./examples/display/sample.gif)](./examples/display/sample.gif)

## io-go, get it?

Download the library:

```bash
go get github.com/zarthus/iogo
```

## Usage

In essence: You'll want to create an `iogo.ReadWriter`, the easiest way is with `iogo.style.NewStdReadWriter`

Refer to `main.go` for a detailed list of examples.

### Examples

Examples are found in [directory ./examples](./examples)

To run an example, you can use 

```
$ go run main.go filename_without_extension
```

To run all examples, you can use

```
$ go run main.go all
```

## Supported Terminal Emulator Versions

The aim (though not the resources) is to support "modern" terminals.


- Modern Microsoft published terminal emulators
  - Specifically "Windows Terminal" on the store, not the old cmd.exe
- Modern versions of popular terminal emulators on *nix, such as GNOME Terminal & Konsole
- MacOS Terminal.app
 
The assumption is made that the user is using a terminal emulator of a relatively new version
  on a relatively modern version of their operating system.

## License

[MIT license](LICENSE)
