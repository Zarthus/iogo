package test

type nullWriter struct {
	writes []string
}

func NewNullWriter() *nullWriter {
	return &nullWriter{
		writes: []string{},
	}
}

func (w *nullWriter) Write(p []byte) (int, error) {
	w.writes = append(w.writes, string(p))
	return 0, nil
}

func (w *nullWriter) Writeln(msg string) (int, error) {
	w.writes = append(w.writes, msg)
	return 0, nil
}

func (w *nullWriter) WriteString(msg string) (int, error) {
	w.writes = append(w.writes, msg)
	return 0, nil
}

func (w nullWriter) Get() []string {
	return w.writes
}

func (w nullWriter) Close() error {
	return nil
}
