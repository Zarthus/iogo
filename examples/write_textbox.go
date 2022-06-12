package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"github.com/zarthus/iogo/v2/pkg/term/col"
)

func WriteTextBox() {
	rw := style.NewStdReadWriter()

	msg := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. " +
		"Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, " +
		"when an unknown printer took a galley of type and scrambled it to make a type specimen book. " +
		"It has survived not only five centuries, but also the leap into electronic typesetting, " +
		"remaining essentially unchanged."

	if err := rw.OutputStyle().TextBox(msg, iogo.Options{}); err != nil {
		panic(err)
	}

	fuchsia, black, blue := col.Fuchsia, col.Black, col.Cyan

	if err := rw.OutputStyle().TextBox(msg, iogo.Options{BgColour: &fuchsia}); err != nil {
		panic(err)
	}

	if err := rw.OutputStyle().TextBox(msg, iogo.Options{FgColour: &fuchsia}); err != nil {
		panic(err)
	}

	if err := rw.OutputStyle().TextBox(msg, iogo.Options{BgColour: &blue, FgColour: &black}); err != nil {
		panic(err)
	}
}
