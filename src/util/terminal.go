package util

import (
	"os"
	"os/exec"
	"time"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayMessageAndWait(message string, seconds int) {
	ClearConsole()
	println(message)
	time.Sleep(time.Duration(seconds) * time.Second)
	ClearConsole()
}

func DisplayMessage(message string) {
	ClearConsole()
	println(message)
}

func DisplayErrorAndWait(message string, seconds int) {
	ClearConsole()
	println(message)
	time.Sleep(time.Duration(seconds) * time.Second)
	ClearConsole()
}
