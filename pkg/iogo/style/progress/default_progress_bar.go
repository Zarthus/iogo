package progress

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
		return
	}

	if bar.current+num >= bar.maximum {
		bar.Finish()
	} else {
		bar.current += num
	}
}

func (bar *defaultProgressBar) SetMaximum(max uint) {
	if bar.locked {
		// TODO: maybe we can avoid panicing, probably impl new error type and return (error)?
		panic("Cannot set maximum when bar is already finished")
	}
	bar.maximum = max
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
