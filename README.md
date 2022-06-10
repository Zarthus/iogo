# iogo 

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/Zarthus/iogo.svg)](https://pkg.go.dev/github.com/Zarthus/iogo/v2)
[![Coverage Status](https://coveralls.io/repos/github/Zarthus/iogo/badge.svg?branch=main)](https://coveralls.io/github/Zarthus/iogo?branch=main)

`iogo` is a helper go library for handling input and output, and stylizing them.

At it's core, iogo is a reader and a writer that doesn't do a whole more than
`fmt.Print` and `fmt.Scanf` with pretty styling options, and if you don't need that,
you probably don't need this library.

Where iogo helps you, is by offering an extensive toolkit for history management,
terminal colours, handling input, and providing styling for output.

Currently, it is not entirely feature complete (support for table outputting,
better term detection & 256-colour support are things I'd really like to include).

[![cast](./examples/display/sample.gif)](./examples/display/sample.gif)

## io-go, get it?

Download the library:

```bash
go get github.com/zarthus/iogo
```

## Usage

In essence: You want to either instantiate a `reader`, a `writer`
or a `io` based on your purpose.

Refer to `main.go` for a detailed example.

### Very basic usage

```go
rw := style.CreateStdReadWriter()

rw.Writer().Writeln("What's your name?") // Or use `name, err := rw.InputStyle().Prompt(...)`
if name, err := rw.Reader().Readln(); err != nil {
    panic(err)
} else {
    rw.Writer().Writeln("Nice to meet you, " + name)
}
```

### Input handling

```go
rw := style.CreateStdReadWriter()

confirmed, err := rw.InputStyle().Confirm("Do you want to go swimming today?", iogo.Options{Default: "y"})
if err != nil {
    panic(err)
}

if confirmed {
    goSwimming()
}
```

### Output handling

```go
rw := style.CreateStdReadWriter()

rw.OutputStyle().Title("Welcome to iogo!")
rw.OutputStyle().Success("You have installed the software correctly!")
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
