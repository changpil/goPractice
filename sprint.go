package main

import (
	"io/ioutil"
	"fmt"
)


func main() {

    dir, err := ioutil.TempDir("/tmp", "BDDHelmCertificate")
    if err != nil {
	    fmt.Println("AAAA")
	    fmt.Println(err)
	    return
    }
    fmt.Println(dir)
}

