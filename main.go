package main

import (
	"flag"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"os"
	"time"
)

var (
	confirmFlag  = flag.Bool("confirm", false, "use a confirmation question")
	selectFlag   = flag.Bool("select", false, "use a multiple-choice selection")
	progressFlag = flag.Bool("progress", false, "shows a progress bar")
	helpFlag     = flag.Bool("help", false, "show help text")
)

const helpText = "iogo " + iogo.Version + "\n\n" +
	"USAGE:\n" +
	"  --help      this help text\n" +
	"  --confirm   use a confirmation question\n" +
	"  --select    use a multiple-choice selection\n" +
	"  --progress  show a progress bar\n\n"

type flags struct {
	confirmFlag  bool
	selectFlag   bool
	progressFlag bool
	helpFlag     bool
}

func main() {
	flag.Parse()
	f := flags{
		confirmFlag:  *confirmFlag,
		selectFlag:   *selectFlag,
		progressFlag: *progressFlag,
		helpFlag:     *helpFlag,
	}

	os.Exit(demo(f))
}

func demo(f flags) int {
	rw := style.CreateReadWriter(os.Stdin, os.Stdout)

	if f.helpFlag {
		rw.Writer().WriteString(helpText)
		return 0
	}
	if f.selectFlag && f.confirmFlag {
		rw.Writer().Writeln("Options --confirm and --select cannot be used in conjunction")
		return 1
	}

	inp, err := readInput(rw, f)

	if err != nil {
		rw.Writer().Writeln("error! " + err.Error())
		panic(err)
	}

	rw.OutputStyle().Success("Output: " + inp)
	return 0
}

func readInput(rw iogo.Iogo, f flags) (string, error) {
	opts := iogo.Options{
		DoNotTrack: true,
	}

	rw.OutputStyle().Title("Welcome to iogo!")
	inStyle := rw.InputStyle()

	if f.progressFlag {
		bar := progress.NewDefaultProgressBar(10)
		descriptor := "Determining which input to offer.."
		fmter := formatter.NewSimpleProgressBarFormatter(&descriptor)
		rw.OutputStyle().Progress(bar, func(progressBar iogo.ProgressBar) {
			progressBar.Advance(1)
			time.Sleep(100 * time.Millisecond)
		}, &fmter)
	}

	if f.confirmFlag {
		n := "n"
		opts.Default = &n

		if confirmed, err := inStyle.Confirm("Do you agree to the terms and conditions?", opts); err != nil {
			return "", err
		} else if confirmed {
			return "You confirmed!", nil
		} else {
			return "You did not confirm.", nil
		}
	} else if f.selectFlag {
		return inStyle.Select("Pick a number:", []string{"one", "two", "three"}, opts)
	} else {
		return inStyle.Prompt("Please insert text:", opts)
	}
}
