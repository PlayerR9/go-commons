package strings

import (
	"bytes"
	"io"
)

type LineBuffer struct {
	lines [][]byte
}

func (lb LineBuffer) String() string {
	return string(bytes.Join(lb.lines, []byte("\n")))
}

func (lb *LineBuffer) Write(data []byte) (int, error) {
	if lb == nil {
		return 0, io.ErrShortWrite
	}

	if len(data) == 0 {
		lb.lines = append(lb.lines, []byte{})
	}

	lines := bytes.Split(data, []byte("\n"))

	for _, line := range lines {
		lb.lines = append(lb.lines, line)
	}

	return len(data), nil
}

func (lb *LineBuffer) AddString(line string) {
	if lb == nil {
		return
	}

	data := []byte(line)

	lines := bytes.Split(data, []byte("\n"))
	lb.lines = append(lb.lines, lines...)
}

func (lb LineBuffer) LinesString() []string {
	lines := make([]string, 0, len(lb.lines))

	for _, line := range lb.lines {
		lines = append(lines, string(line))
	}

	return lines
}

func (lb LineBuffer) Lines() [][]byte {
	lines := make([][]byte, 0, len(lb.lines))

	for _, line := range lb.lines {
		row := make([]byte, len(line))
		copy(row, line)

		lines = append(lines, row)
	}

	return lines
}

func (lb *LineBuffer) Reset() {
	if lb == nil {
		return
	}

	if len(lb.lines) > 0 {
		for i := 0; i < len(lb.lines); i++ {
			for j := 0; j < len(lb.lines[i]); j++ {
				lb.lines[i][j] = 0
			}

			lb.lines[i] = lb.lines[i][:0]
			lb.lines[i] = nil
		}

		lb.lines = lb.lines[:0]
	}
}
