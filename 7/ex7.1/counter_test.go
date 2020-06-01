package ex7_1

import (
	"testing"
)

var input = "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"

func TestWordsCounter_Write(t *testing.T) {

	var w WordsCounter
	w.Write(input)
	t.Logf("word count = %d\n", w)
	if w != 32 {
		t.Fail()
	}

}

func TestLinesCounter_Write(t *testing.T) {
	var l LinesCounter
	l.Write(input)
	t.Logf("line count = %d\n", l)
	if l != 3 {
		t.Fail()
	}
}
