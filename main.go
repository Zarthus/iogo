package main

import (
	"flag"
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"os"
)

var confirmFlag = flag.Bool("confirm", false, "use a confirmation question")
var selectFlag = flag.Bool("select", false, "use a multiple-choice selection")
var helpFlag = flag.Bool("help", false, "show help text")

type flags struct {
	confirmFlag bool
	selectFlag  bool
	helpFlag    bool
}

func main() {
	flag.Parse()
	f := flags{
		confirmFlag: *confirmFlag,
		selectFlag:  *selectFlag,
		helpFlag:    *helpFlag,
	}
	if ok := demo(f); !ok {
		os.Exit(1)
	}
}

func demo(f flags) bool {
	rw := style.CreateDefaultReadWriter()
	if f.helpFlag {
		showHelp(rw)
	} else if f.confirmFlag && f.selectFlag {
		rw.Writer().Writeln("Options --confirm and --select cannot be used in conjunction")
		return false
	} else if in, err := readInput(rw, f); err != nil {
		rw.Writer().Writeln("Error! " + err.Error())
		return false
	} else {
		rw.Style().Output().Success("Output: " + in)
	}
	return true
}

func readInput(rw iogo.ReadWriter, f flags) (string, error) {
	opts := &iogo.Options{
		DoNotTrack: true,
	}
	rw.Style().Output().Title("Welcome to iogo!")
	inStyle := rw.Style().Input()
	if f.confirmFlag {
		opts.Default = "n"
		if ok, err := inStyle.Confirm("Do you agree to the terms and conditions?", opts); err != nil {
			return "", err
		} else if ok {
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

func showHelp(rw iogo.ReadWriter) {
	rw.Writer().Writeln(fmt.Sprintf(`iogo %s

USAGE:
  --help      this help text
  --confirm   use a confirmation question
  --select    use a multiple-choice selection
`, iogo.Version))
}
