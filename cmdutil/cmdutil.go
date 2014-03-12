// This package provides a set of utilities useful for command line interaction.
package cmdutil

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ReadLine prompts the given message and reads a line from the command line.
// It returns the line and an error if that fails.
func ReadLine(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(line)), nil
}

// ReadSilentLine prompts the given message and reads a line from the command
// line.
// The line won't be seen on the command line, this is very useful for
// passwords and other secret information.
// It returns the line and an error if that fails.
func ReadSilentLine(prompt string) (string, error) {
	silence()
	defer func() {
		unsilence()
		// we add a blank line after unsilencing.
		fmt.Println()
	}()
	return ReadLine(prompt)
}

func silence() {
	run(exec.Command("stty", "-echo"))
}

func unsilence() {
	run(exec.Command("stty", "echo"))
}

func run(command *exec.Cmd) {
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Run()
}
