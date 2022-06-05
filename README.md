# iogo 

`iogo` is a helper go library for handling input and output, and stylizing them.

At it's core, iogo is a reader and a writer that doesn't do a whole more than
`fmt.Print` and `fmt.Scanf`, and if that's all that you need, you probably don't
need this library - or you're just dealing with very little input.

Where iogo helps you, however, is offering an extensive toolkit for history management, 
terminal colours, handling input, and providing styling for output.

Currently, it is not entirely feature complete (support for table outputting, 
better term detection & 256-colour support are things I'd really like to include).

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
io := style.CreateDefaultIo()

io.GetWriter().WriteLine("What's your name?")
name, err := io.GetReader().ReadLine()
if err != nil {
    panic(err)
}
io.GetWriter().WriteLine("Nice to meet you, " + name)
```

### Input handling

```go
io := style.CreateDefaultIo()

confirmed, err := io.Style().Input().Confirm("Do you want to go swimming today?", iogo.Options{Default: "y"})
if err != nil {
    panic(err)
}

if confirmed {
    goSwimming()
}
```

### Output handling

```go
io := style.CreateDefaultIo()

io.Style().Output().Title("Welcome to iogo!")
io.Style().Output().Success("You have installed the software correctly!")
```

## License

[MIT license](LICENSE)
