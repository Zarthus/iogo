package iogo

type Reader interface {
	ReadLine(opts *Options) (string, error)
	Reset()
}

type Writer interface {
	Write(msg string)
	Writeln(msg string)
}

type ReaderStyle interface {
	Prompt(prompt string, opts *Options) (string, error)
	Confirm(prompt string, opts *Options) (bool, error)
	Select(prompt string, valid []string, opts *Options) (string, error)
}

type WriterStyle interface {
	Title(msg string)
	Section(msg string)

	Block(msg string, opts *Options)

	Info(msg string)
	Success(msg string)
	Warning(msg string)
	Error(msg string)
}

type Style interface {
	Input() ReaderStyle
	Output() WriterStyle
}

type ReadWriter interface {
	Reader() Reader
	Writer() Writer
	Style() Style
}

type HistoryTracker interface {
	Get() []string
	Track(item string)
	Untrack(item string)
	Save()
	Reset()
}
