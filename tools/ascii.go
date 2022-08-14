package main

import (
	"fmt"
	"os/exec"
)

func main() {
	Colors()
}

func Table() {
	for i := 0; i < 127; i++ {
		fmt.Printf("%d:\t%s\n", i, string(rune(i)))
	}
}

func Colors() {
	exec.Command("color").Run()
	for i := 0; i < 40; i++ {
		fmt.Printf("%d: \033[%dmHello world\033[0m\n", i, i)
	}
}
