package main

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"os"
	"time"
)

func main() {
	exitCode := demo(os.Args)

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func demo(args []string) int {
	rw := style.CreateDefaultReadWriter()
	sel, conf, prog, exitcode := parseOpts(rw, args)
	if exitcode >= 0 {
		return exitcode
	}

	inp, err := readInput(rw, sel, conf, prog)

	if err != nil {
		rw.Writer().WriteLine("error! " + err.Error())
		panic(err)
	}

	rw.Style().Output().Success("Output: " + inp)
	return 0
}

func parseOpts(rw iogo.Iogo, args []string) (bool, bool, bool, int) {
	sel, conf, prog := false, false, false
	exitcode := -1

	for _, arg := range args {
		if arg == "--select" {
			sel = true
		}
		if arg == "--confirm" {
			conf = true
		}
		if arg == "--progress" {
			prog = true
		}
		if arg == "--help" {
			exitcode = 0
			rw.Writer().Write(
				"iogo " + iogo.Version + "\n\n" +
					"USAGE:\n" +
					"  --help      this help text\n" +
					"  --confirm   use a confirmation question\n" +
					"  --select    use a multiple-choice selection\n" +
					"  --progress  show a progress bar\n\n",
			)
		}
	}

	if sel && conf {
		rw.Writer().WriteLine("options --confirm and --select cannot be used in conjunction")
		return sel, conf, prog, 1
	}

	return sel, conf, prog, exitcode
}

func readInput(io iogo.Iogo, sel bool, conf bool, prog bool) (string, error) {
	var inp string
	var err error

	options := iogo.Options{
		DoNotTrack: true,
	}

	io.Style().Output().Title("Welcome to iogo!")
	inpstyle := io.Style().Input()
	if !sel && !conf {
		inp, err = inpstyle.Prompt("Please insert text:", options)
	} else if sel {
		inp, err = inpstyle.Select("Pick a number:", []string{"one", "two", "three"}, options)
	} else if conf {
		options.Default = "n"
		conf, err = inpstyle.Confirm("Do you agree to the terms and conditions?", options)

		if conf {
			inp = "You confirmed!"
		} else {
			inp = "You did not confirm."
		}
	}

	if prog {
		bar := progress.NewDefaultProgressBar(10)
		io.Style().Output().Progress(bar, func(progressBar iogo.ProgressBar) {
			progressBar.Advance(1)
			time.Sleep(100 * time.Millisecond)
		}, nil)
	}

	return inp, err
}
