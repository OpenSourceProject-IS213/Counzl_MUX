package cmd

import (
	"time"
	"os/exec"
)

func Loading() {

	for i := 0; i < 5; i++ {
		time.Sleep(220 * time.Millisecond)
		print("\r|")
		time.Sleep(220 * time.Millisecond)
		print("\r/")
		time.Sleep(220 * time.Millisecond)
		print("\r-")
		time.Sleep(220 * time.Millisecond)
		print("\r\\")
		time.Sleep(220 * time.Millisecond)
		print("\r|")
		time.Sleep(220 * time.Millisecond)
		print("\r/")
		time.Sleep(220 * time.Millisecond)
		print("\r-")
		time.Sleep(220 * time.Millisecond)
		print("\r\\")
	}
	time.Sleep(220 * time.Millisecond)
	print("\r")
}

func clearScreen(){
	exec.Command("clear")
}

