package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func timeRecorder(fname string) func() {
	now := time.Now().UnixNano()
	return func() {
		fmt.Printf("function %s consume time = %v ns\n", fname, time.Now().UnixNano()-now)
	}
}

func echo1() {
	funcPtr, _, _, ok := runtime.Caller(0)
	var funcName string
	if ok {
		funcName = runtime.FuncForPC(funcPtr).Name()
	}

	defer timeRecorder(funcName)()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	funcPtr, _, _, ok := runtime.Caller(0)
	var funcName string
	if ok {
		funcName = runtime.FuncForPC(funcPtr).Name()
	}

	defer timeRecorder(funcName)()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	funcPtr, _, _, ok := runtime.Caller(0)
	var funcName string
	if ok {
		funcName = runtime.FuncForPC(funcPtr).Name()
	}

	defer timeRecorder(funcName)()
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	// os.Args[1:] = {1,2,3,4}
	echo1() // 17000ns
	echo2() // 2000ns
	echo3() // 1000ns
}
