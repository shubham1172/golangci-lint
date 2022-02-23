package testdata

import (
	"bytes"
	"os/exec"
)

func GokartRunCommand(userInput string) error {
	cmd := exec.Command(userInput)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run() // ERROR ""
	if err != nil {
		return err
	}
	return nil
}
