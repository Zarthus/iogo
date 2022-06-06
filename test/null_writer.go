package test

type nullWriter struct {
	writes []string
}

func NewNullWriter() *nullWriter {
	return &nullWriter{
		writes: []string{},
	}
}

func (writer *nullWriter) Write(message string) {
	writer.writes = append(writer.writes, message)
}

func (writer *nullWriter) Writeln(message string) {
	writer.writes = append(writer.writes, message)
}

func (writer nullWriter) Get() []string {
	return writer.writes
}
