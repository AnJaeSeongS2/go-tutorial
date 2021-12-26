package methods

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct {
	readIndex int
	msg       []byte
}

func NewMyReader(s string) *MyReader {
	r := MyReader{}
	r.msg = []byte(s)
	return &r
}

func (m *MyReader) Read(b []byte) (n int, err error) {
	if m.readIndex >= len(m.msg) {
		return 0, io.EOF
	}

	readSize := copy(b, m.msg[m.readIndex:])
	m.readIndex += readSize
	return readSize, nil
}

func TourExerciseReader() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	reader.Validate(NewMyReader("A"))
}
