package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetChanges() string {
	cmd := exec.Command("git", "status", "--short")
	path, _ := GetFromRoot(".")
	cmd.Dir = path

	output, _ := cmd.CombinedOutput()

	return string(output)
}

func CommitChanges(message string, path string) {
	err := exec.Command("git", "add", path).Run()
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := exec.Command("git", "commit", "-sm", message).CombinedOutput()

	if err2 != nil {
		fmt.Println(err2)
	}
}

func SetEnvFromFile() {
	env, _ := GetFromRoot(".env")
	data, _ := GetFileData(env)
	vars := strings.SplitSeq(data, "\n")
	for variable := range vars {
		spt := strings.Split(variable, "=")
		os.Setenv(spt[0], spt[1])
	}
}

func WriteLog(line string) {
	file := "/tmp/go-api-log.txt"
	data, _ := GetFileData(file)
	data += time.Now().Format("2006-01-02 15:04:05") + " " + line + "\n"
	SaveFile(file, data)
}

func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func DecodeJson(data io.Reader, variable any) error {
	decoder := json.NewDecoder(data)

	if err := decoder.Decode(&variable); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func EncodeJson(data io.Writer, variable any) error {
	decoder := json.NewEncoder(data)

	if err := decoder.Encode(&variable); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetDate() (string, string) {
	now := time.Now()
	before, after, _ := strings.Cut(now.Format(time.RFC3339Nano), "T")

	return before, after
}

func FileExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
