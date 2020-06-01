package ex7_2

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("Spicy jalapeno pastrami")
	c.Write(data)
	if *n != int64(len(data)) {
		t.Logf("%d != %d", n, len(data))
		t.Fail()
	}
}
