package ex7_1

import (
	"bufio"
	"strings"
)

type WordsCounter int

func (c *WordsCounter) Write(in string) (int, error) {
	r := strings.NewReader(in)
	scanner := bufio.NewScanner(r)
	// set split function
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return int(*c), scanner.Err()
}

type LinesCounter int

func (l *LinesCounter) Write(in string) (int, error) {
	r := strings.NewReader(in)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}
	return int(*l), scanner.Err()
}
