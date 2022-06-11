package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
)

func WriteBlocks() {
	rw := style.NewStdReadWriter()

	rw.OutputStyle().Block("Unstyled Block", iogo.Options{})
	rw.OutputStyle().Info("Info Block")
	rw.OutputStyle().Success("Success Block")
	rw.OutputStyle().Warning("Warning Block")
	rw.OutputStyle().Error("Error Block")
}
