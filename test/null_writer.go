package test

type nullWriter struct {
	writes []string
}

func NewNullWriter() *nullWriter {
	return &nullWriter{
		writes: []string{},
	}
}

func (writer *nullWriter) Write(message string) (int, error) {
	writer.writes = append(writer.writes, message)
	return 0, nil
}

func (writer *nullWriter) Writeln(message string) (int, error) {
	writer.writes = append(writer.writes, message)
	return 0, nil
}

func (writer nullWriter) Get() []string {
	return writer.writes
}
