package test

import (
	"fmt"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/dollarkillerx/processes"
	"os"
	"path"
	"testing"
)

func TestP2p(t *testing.T) {
	p := map[string][]string{}
	p["aaaa"] = append(p["aaaa"], "asdasd")
	p["aaaa"] = append(p["aaaa"], "asdasd1")
	fmt.Println(p)
}

func TestDir(t *testing.T) {
	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//index := strings.LastIndex(path, string(os.PathSeparator))
	//path = path[:index]
	//
	//command := exec.Command("pwd")
	//command.Dir = path
	//out, err := command.CombinedOutput()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(out))
	//
	//cmd := exec.Command("cd /home/wangy/mvalley/new_2022/dot/dot")
	//fmt.Println(cmd)
	//command = exec.Command("pwd")
	//out, err = command.CombinedOutput()
	//if err != nil {
	//	panic(err)
	//}

	//runCommand, err := processes.RunCommand("cd /home/wangy/mvalley/new_2022/dot/dot")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(runCommand)
	//
	//runCommand, err = processes.RunCommand("pwd")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(runCommand)
	//
	//runCommand, err = processes.RunCommand("bash cmd/dot/build_image.sh")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(runCommand)
	//fmt.Println(string(out))

	cmd := processes.NewExecLinux()
	exec, err := cmd.Exec("cd /home/wangy/mvalley/new_2022/dot/dot")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(exec)

	//cmd.Exec("cd ../../")

	exec, err = cmd.Exec("bash cmd/dot/build_image.sh")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(exec)
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

func TestPPx(t *testing.T) {
	cmd := processes.NewExecLinux()
	_, err := cmd.Exec("cd /home/wangy/workspace/2022/fireworks/test")
	if err != nil {
		panic(err)
		return
	}

	path2, err := cmd.Exec("pwd")
	if err != nil {
		return
	}
	fmt.Println(path2)
	path2 = path.Join(path2, "xxsadsa", "asdsadsa")
	fmt.Println(path2)

	err = os.MkdirAll(path2, 00777)
	if err != nil {
		panic(err)
	}
}
