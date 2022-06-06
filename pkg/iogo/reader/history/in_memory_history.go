package history

type inMemoryHistory struct {
	history []string
}

func NewHistoryTracker(history []string) *inMemoryHistory {
	return &inMemoryHistory{history: history}
}

func (hist inMemoryHistory) Get() []string {
	return hist.history
}

func (hist *inMemoryHistory) Track(item string) {
	hist.history = append(hist.history, item)
}

func (hist *inMemoryHistory) Untrack(item string) bool {
	var index int

	for key, value := range hist.history {
		if value == item {
			index = key
		}
	}

	if nil == &index {
		return false
	}

	hist.history = append(hist.history[:index], hist.history[index+1:]...)
	return true
}

func (hist inMemoryHistory) Save() bool {
	return false
}

func (hist *inMemoryHistory) Reset() {
	hist.history = []string{}
}
