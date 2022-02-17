package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dollarkillerx/fireworks/internal/pkg/utils"
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

func TestZip(t *testing.T) {
	open, err := os.Open("../internal")
	if err != nil {
		panic(err)
	}

	err = utils.Compress([]*os.File{open}, "data.zip")
	if err != nil {
		panic(err)
	}

	os.MkdirAll("data", 00755)
	err = utils.DeCompress("data.zip", "data")
	if err != nil {
		panic(err)
	}
}
