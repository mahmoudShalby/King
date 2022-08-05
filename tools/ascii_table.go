package main

import "fmt"

func main() {
	for i := 0; i < 127; i++ {
		fmt.Printf("%d:\t%s\n", i, string(rune(i)))
	}
}
