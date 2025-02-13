package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func MergeDump(filename string) error {
	name := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	output, _ := exec.Command("pwd").Output()
	pwd := strings.TrimSpace(string(output))

	cmdArgs := []string{
		"-u", username,
		"-p" + password,
		name,
		"-e", fmt.Sprintf("source %s/dump/uncompressed/%s.sql", pwd, filename),
	}

	_, err := exec.Command("mysql", cmdArgs...).Output()
	if err != nil {
		return err
	}

	return nil
}
