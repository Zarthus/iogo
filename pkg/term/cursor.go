package term

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
)

const (
	CursorGet      ControlSequence = "\033[6n"
	cursorPos      ControlSequence = "\033[%d;%dR"
	CursorSet      ControlSequence = "\033[%d;%dH" // position the cursor at Line N and Column N
	CursorSetAlt   ControlSequence = "\033[%d;%df" // Alternative way of setting cursor
	cursorUp       ControlSequence = "\033[%dA"
	cursorDown     ControlSequence = "\033[%dB"
	cursorForward  ControlSequence = "\033[%dC"
	cursorBackward ControlSequence = "\033[%dD"
	CursorShow     ControlSequence = "\033[?25h"
	CursorHide     ControlSequence = "\033[?25l"
)

type CursorInstruction struct {
	Writer iogo.Writer

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

func (c *CursorInstruction) Write(msg string) *CursorInstruction {
	return c.handle(msg)
}

func (c *CursorInstruction) handle(msg string) *CursorInstruction {
	_, err := c.Writer.WriteString(msg)
	c.handleErr(&err)
	return c
}

func (c *CursorInstruction) handleErr(err *error) *CursorInstruction {
	if err != nil {
		if c.Err == nil {
			c.Err = *err
		}
	}
	return c
}
