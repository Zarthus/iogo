package iogo

type Reader interface {
	ReadLine(options Options) (string, error)
	Reset()
}

type Writer interface {
	Write(message string)
	WriteLine(message string)
}

type ReaderStyle interface {
	Prompt(prompt string, options Options) (string, error)
	Confirm(prompt string, options Options) (bool, error)
	Select(prompt string, valid []string, options Options) (string, error)
}

type WriterStyle interface {
	Title(message string)
	Section(message string)

	Block(message string, options Options)

	Info(message string)
	Success(message string)
	Warning(message string)
	Error(message string)
}

type Style interface {
	Input() ReaderStyle
	Output() WriterStyle
}

type Io interface {
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
