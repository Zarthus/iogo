package term

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
)

const (
	CursorSave     ControlSequence = "\033[s"
	CursorRestore  ControlSequence = "\033[u"
	cursorUp       ControlSequence = "\033[%dA"
	cursorDown     ControlSequence = "\033[%dB"
	cursorForward  ControlSequence = "\033[%dC"
	cursorBackward ControlSequence = "\033[%dD"
)

type CursorInstruction struct {
	writer iogo.Writer
}

func (c *CursorInstruction) Up(num int) *CursorInstruction {
	c.writer.Write(fmt.Sprintf(string(cursorUp), num))
	return c
}

func (c *CursorInstruction) Down(num int) *CursorInstruction {
	c.writer.Write(fmt.Sprintf(string(cursorDown), num))
	return c
}

func (c *CursorInstruction) Forward(num int) *CursorInstruction {
	c.writer.Write(fmt.Sprintf(string(cursorForward), num))
	return c
}

func (c *CursorInstruction) Backward(num int) *CursorInstruction {
	c.writer.Write(fmt.Sprintf(string(cursorBackward), num))
	return c
}

func (c *CursorInstruction) SavePosition() *CursorInstruction {
	c.writer.Write(string(CursorSave))
	return c
}

func (c *CursorInstruction) RestorePosition() *CursorInstruction {
	c.writer.Write(string(CursorRestore))
	return c
}

func (c *CursorInstruction) Write(msg string) *CursorInstruction {
	c.writer.Write(msg)
	return c
}
