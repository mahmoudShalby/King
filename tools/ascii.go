package main

import "fmt"

func main() {
	Table()
}

func Table() {
	for i := 0; i < 127; i++ {
		fmt.Printf("%d:\t%s\n", i, string(rune(i)))
	}
}

func Colors() {
	for i := 0; i < 40; i++ {
		fmt.Printf("%d: \x1b[%dmHello world\x1b[0m\n", i, i)
	}
}
