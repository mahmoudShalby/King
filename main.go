package main

import (
	"fmt"
	"io/ioutil"
	"king/parser"
	"os"
)

func main() {
	text := readFile("a.king")
	if len(text) != 0 {
		var p parser.Parser
		p.Init(text)
	}
}

func readFile(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(text)
}
