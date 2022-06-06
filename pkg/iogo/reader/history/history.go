package history

type tracker struct {
	history []string
}

func NewHistoryTracker(history []string) *tracker {
	return &tracker{history}
}

func (t tracker) Get() []string {
	return t.history
}

func (t *tracker) Track(item string) {
	t.history = append(t.history, item)
}

func (t tracker) Untrack(item string) {
	panic("implement me") // TODO: implement
}

func (t tracker) Save() {
	panic("implement me") // TODO: implement
}

func (t *tracker) Reset() {
	t.history = []string{}
}
