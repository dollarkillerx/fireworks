package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestDir(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]

	command := exec.Command("pwd")
	command.Dir = path
	out, err := command.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
