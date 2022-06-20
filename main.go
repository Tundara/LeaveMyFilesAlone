package main

import (
	"LetMyFiles/gui"
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		gui.StartGui()
	} else if runtime.GOOS == "linux" {
		fmt.Println("Linux")
	}
}
