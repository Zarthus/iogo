package progress

import (
	"errors"
)

type defaultProgressBar struct {
	current uint
	maximum uint
	locked  bool
}

func NewDefaultProgressBar(maximum uint) *defaultProgressBar {
	return &defaultProgressBar{
		current: 0,
		maximum: maximum,
		locked:  false,
	}
}

func (bar *defaultProgressBar) Advance(num uint) {
	if bar.locked {
		// If you want to explicitly only advance when not locked to ensure input is not lost, you can use IsFinished()
		return
	}

	if bar.current+num >= bar.maximum {
		bar.Finish()
	} else {
		bar.current += num
	}
}

func (bar *defaultProgressBar) SetMaximum(max uint) error {
	if bar.locked {
		return errors.New("cannot set maximum when bar is finished")
	}
	bar.maximum = max
	return nil
}

func (bar defaultProgressBar) Current() uint {
	return bar.current
}

func (bar defaultProgressBar) Maximum() uint {
	return bar.maximum
}

func (bar *defaultProgressBar) Finish() {
	if bar.locked {
		return
	}
	bar.locked = true
	bar.current = bar.maximum
}

func (bar defaultProgressBar) IsFinished() bool {
	if bar.locked {
		return true
	}

	if bar.current >= bar.maximum {
		bar.Finish()
		return true
	}

	return false
}
