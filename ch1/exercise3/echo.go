package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//? The prevous is always much slower than the latter
	//? Even if I change the order of the stdJoin() and inEfficient()
	//? WTF??? Does it has any relation with I'm running big data progorm that occupied most of the CPU?

	start := time.Now()
	inEfficient()
	fmt.Println("inefficient in book:", time.Since(start).String())
	stdStart := time.Now()
	stdJoin()
	fmt.Println("std strings.Join():", time.Since(stdStart).String())
}

func stdJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func inEfficient() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
