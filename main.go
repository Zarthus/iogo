package main

import (
	"fmt"
	"github.com/zarthus/iogo/v2/examples"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"os"
	"time"
)

type mapping struct {
	Name     string
	Runnable func()
}

func main() {
	rw := style.NewStdReadWriter()
	mappings := loadMappings()

	if len(os.Args) > 1 {
		selection := os.Args[1]
		exec(selection, mappings)
		return
	}

	selection, err := rw.InputStyle().Select("Please select a program to run", selectables(mappings), iogo.Options{})

	if err != nil {
		rw.OutputStyle().Error(err.Error())
		return
	}

	rw.OutputStyle().Info("Starting " + selection)
	exec(selection, mappings)
}

func loadMappings() []mapping {
	return []mapping{
		{"all", multiple},
		{"progress_bars_multiple", examples.ProgressBarsMultiple},
		{"progress_bars_single", examples.ProgressBarsSingle},
		{"read_confirm", examples.ReadConfirm},
		{"read_name", examples.ReadName},
		{"write_blocks", examples.WriteBlocks},
		{"write_style", examples.WriteStyle},
		{"write_textbox", examples.WriteTextBox},
	}
}

func multiple() {
	m := loadMappings()
	mappings := m[1:]

	for _, mapped := range mappings {
		fmt.Printf("\n\n%s\n", mapped.Name)
		mapped.Runnable()
		time.Sleep(time.Second * 2)
	}
}

func selectables(mappings []mapping) []string {
	var s []string
	for _, mappable := range mappings {
		s = append(s, mappable.Name)
	}
	return s
}

func exec(selection string, mappings []mapping) {
	for _, mappable := range mappings {
		if mappable.Name == selection {
			mappable.Runnable()
			return
		}
	}
	_ = fmt.Errorf("could not find program %s", selection)
	os.Exit(1)
}
