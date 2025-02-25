package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getPwd() string {
	output, _ := exec.Command("pwd").Output()
	pwd := strings.TrimSpace(string(output))

	return pwd
}

func MergeDumpMySQL(filename string) error {
	name := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")

	pwd := getPwd()

	cmdArgs := []string{
		"-u", username,
		"-p" + password,
		"-h", host,
		name,
		"-e", fmt.Sprintf("source %s/dump/uncompressed/%s.sql", pwd, filename),
	}

	_, err := exec.Command("mysql", cmdArgs...).Output()
	if err != nil {
		return err
	}

	return nil
}

func MergeDumpHadoop(filename string) error {
	pwd := getPwd()

	cmdArgs := []string{
		"dfs",
		"-put",
		fmt.Sprintf("%s/dump/uncompressed/%s.sql", pwd, filename),
		"/sensors/",
	}
	_, err := exec.Command("hdfs", cmdArgs...).Output()
	if err != nil {
		return err
	}

	return nil
}
