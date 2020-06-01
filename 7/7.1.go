package main

import (
	"bufio"
	"fmt"
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

func main() {
	input := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"

	var w WordsCounter
	w.Write(input)
	fmt.Println(w)

	var l LinesCounter
	l.Write(input)
	fmt.Println(l)

}
