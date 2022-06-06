package iogo

// Reader is an interface that defines minimistically how to read from somewhere (that is a terminal/tty/some form of stdin)
type Reader interface {
	// ReadLine reads input into a string.
	ReadLine(options Options) (string, error)

	// Reset the state of the Reader, this clears any history and state.
	Reset()
}

// Writer is an interface that defines minimistically how to write to something (that is a terminal/tty/file/some form of stdin)
type Writer interface {
	// Write to the output source of the Writer, without ensuring CRLF or LF at the end.
	Write(message string)
	// WriteLine writes to the output source of the Writer, ensuring CRLF or LF at the end.
	WriteLine(message string)
}

// ReaderStyle is a helpful subset of helper methods that help receive input in a desired format.
// Despite being largely dependent on the Reader, it also writes output in many cases.
type ReaderStyle interface {
	// Prompt the input source of the Reader a free-style question
	Prompt(prompt string, options Options) (string, error)
	// Confirm the input source of the Reader a yes or no question
	Confirm(prompt string, options Options) (bool, error)
	// Select a multiple-choice input to the input source of the Reader, ensuring the return value is one
	// of the items given in valid
	Select(prompt string, valid []string, options Options) (string, error)
}

// WriterStyle is a helpful subset of helper methods that help format output in a desired format.
type WriterStyle interface {
	// Title writes a title to the output of the Writer, by default,
	// it looks something like this, based on implementation and term support:
	//
	// message
	// =======
	//
	Title(message string)
	// Section writes a section to the output of the Writer, by default,
	// it looks something like this, based on implementation and term support:
	// message
	// =======
	Section(message string)

	// Block writes a block to the output of the Writer, by default,
	// it looks something like this, based on implementation and term support:
	//
	//   message
	//
	Block(message string, options Options)

	// Info is an alias for Block with appropriate colour (if supported), or a sane fallback
	Info(message string)
	// Success is an alias for Block with appropriate colour (if supported), or a sane fallback
	Success(message string)
	// Warning is an alias for Block with appropriate colour (if supported), or a sane fallback
	Warning(message string)
	// Error is an alias for Block with appropriate colour (if supported), or a sane fallback
	Error(message string)

	// Progress Renders starts the Progress bar, and loops the runnable until the bar is finished.
	// It is expected that the inside of the runnable calls ProgressBar.Advance to ensure the bar finishes
	// at some point. Therefore, caution: Your program can run into an infinite loop on misuse.
	Progress(bar ProgressBar, runnable func(bar ProgressBar), barFormatter *ProgressBarFormatter)
}

// Style is a combination object that merges ReaderStyle and WriterStyle in one coherent structure, while also offering
// the toolkit for anything that needs "both" in and output or doesn't really strongly fit in one category.
type Style interface {
	Input() ReaderStyle
	Output() WriterStyle
}

// Iogo is the combination of input (Reader) and output (Writer) offering the bare essentials,
// and Style which adds the extra flavour you actually need to make things pretty and awesome.
type Iogo interface {
	Reader() Reader
	Writer() Writer
	Style() Style
}

// ProgressBar helps render progress on some task
type ProgressBar interface {
	// Advance the progress bar by num, or up until Maximum is reached.
	Advance(num uint)
	// SetMaximum increases or decreases the maximum value the progress bar can take (even while it is running).
	// Behaviour of SetMaximum may vary by implementation of the ProgressBar based on if the Current progression
	// exceeds the new Maximum value
	SetMaximum(max uint)

	// Current returns the current number of the progress
	Current() uint
	// Maximum returns the maximum tasks before the ProgressBar is finished.
	Maximum() uint

	// Finish updates the Current value to the Maximum
	Finish()
	// IsFinished returns if the Current value matches the Maximum value
	IsFinished() bool
}

// ProgressBarFormatter formats the progress bar into some style
type ProgressBarFormatter interface {
	// Format a progress bar into a readable string
	Format(bar ProgressBar) string
}

// HistoryTracker keeps track of Reader input to later be reused for e.g. autocompletion
type HistoryTracker interface {
	// Get the current list of stored history
	Get() []string
	// Track a new item to history
	Track(item string)
	// Untrack removes an existing item from history
	// Returns boolean true if item is found and removed, false otherwise
	Untrack(item string) bool
	// Save the current history for later re-use, if the current implementation supports saving.
	// Returns a bool true if saving was successful, or false if it wasn't.
	Save() bool
	// Reset the current history back to a clean slate.
	Reset()
}
