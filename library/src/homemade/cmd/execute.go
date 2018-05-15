package cmd

import (
	"os/exec"
	"log"
)

func Exe_cmd(cmd string) []byte {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
