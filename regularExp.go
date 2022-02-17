package main
import (
	"fmt"
	"regexp"
	"strconv"
	"os/exec"
	"errors"
	"strings"
)


func main() {
   //ss3 := `version.BuildInfo{Version:"v3.0", GitCommit:"", GitTreeState:"", GoVersion:"go1.13.11"}`
   //ss2 := `Client: &version.Version{SemVer:"v2.16.1", GitCommit:"bbdfe5e7803a12bbdf97e94cd847859890cf4050", GitTreeState:"clean"}
   //       Error: could not find tiller`
   
   cmd := "helm2 version"
   output, _ := ShellExec(cmd)
   
   fmt.Println(output)
   target := ":\"v" + strconv.Itoa(3)
   match, _ :=regexp.MatchString(target, output)
   if match {
	fmt.Println("AAAAA")
   } else{
	 fmt.Println("BBBBB")
   }

   fmt.Println(match)
}


func ShellExec(command string) (string, error) {
	normalized_command := strings.Join(strings.Fields(strings.TrimSpace(command)), " ")
    args := strings.Split(normalized_command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd2Output, err := cmd.CombinedOutput()
	if err != nil {
		return string(cmd2Output), errors.New(string(cmd2Output))
	}
	return string(cmd2Output), nil
}
