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
	// Err holds the first retrieved error since the lifecycle of this instance
	// it is likely that if any instruction failed, most after would have as well.
	Err error
}

func (c *CursorInstruction) Up(num int) *CursorInstruction {
	return c.handle(fmt.Sprintf(string(cursorUp), num))
}

func (c *CursorInstruction) Down(num int) *CursorInstruction {
	return c.handle(fmt.Sprintf(string(cursorDown), num))
}

func (c *CursorInstruction) Forward(num int) *CursorInstruction {
	return c.handle(fmt.Sprintf(string(cursorForward), num))
}

func (c *CursorInstruction) Backward(num int) *CursorInstruction {
	return c.handle(fmt.Sprintf(string(cursorBackward), num))
}

func (c *CursorInstruction) SavePosition() *CursorInstruction {
	return c.handle(string(CursorSave))
}

func (c *CursorInstruction) RestorePosition() *CursorInstruction {
	return c.handle(string(CursorRestore))
}

func (c *CursorInstruction) Write(msg string) *CursorInstruction {
	return c.handle(msg)
}

func (c *CursorInstruction) handle(msg string) *CursorInstruction {
	_, err := c.writer.WriteString(msg)
	if err != nil {
		if c.Err == nil {
			c.Err = err
		}
	}
	return c
}
