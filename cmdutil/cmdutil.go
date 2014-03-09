package cmdutil

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func ReadLine() (string, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(line)), nil
}

func Silence() {
	runCommand(exec.Command("stty", "-echo"))
}

func Unsilence() {
	runCommand(exec.Command("stty", "echo"))
}

func runCommand(command *exec.Cmd) {
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Run()
}
