package main

import (
       "errors"
	"os/exec"
        "fmt"
        "strings"
)


func main() {

    cmd := "openssl genrsa -out ./ca.key.pem 4096"
    Exec(cmd)
    cmd = "openssl req -key ca.key.pem -new -x509 -days 7300 -sha256 -out ca.cert.pem -extensions v3_ca"
    Exec(cmd)
    fmt.Println(cmd)
}

func Exec(command string) (string, error) {
    normalized_command := strings.Join(strings.Fields(strings.TrimSpace(command)), " ")
    args := strings.Split(normalized_command, " ")
    cmd := exec.Command(args[0], args[1:]...)
    cmd2Output, err := cmd.CombinedOutput()
    if err != nil {
        return string(cmd2Output), errors.New(string(cmd2Output))
    }
    return string(cmd2Output), nil
}
