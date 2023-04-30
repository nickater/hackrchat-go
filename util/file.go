package util

import (
	"io"
	"os"
)

func SaveToConfigFile(data string) {
	homeDir := os.Getenv("HOME")
	f, err := os.Create(homeDir + "/.hackrchat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, data+"\n")
	if err != nil {
		panic(err)
	}
}

func SaveUsernameToConfigFile(username string) {
	SaveToConfigFile("USERNAME=" + username)
}
