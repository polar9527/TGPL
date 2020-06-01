package ex7_2

import (
	"io"
)

type CountWriter struct {
	Writer io.Writer
	Count  int64
}

func (c *CountWriter) Write(in []byte) (n int, err error) {
	n, err = c.Writer.Write(in)
	c.Count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountWriter{
		Writer: w,
		Count:  0,
	}
	return cw, &cw.Count
}
