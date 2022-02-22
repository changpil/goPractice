package main

import  (
	"encoding/json"
	"fmt"
)

//s := `version.BuildInfo{Version:"v3.0", GitCommit:"", GitTreeState:"", GoVersion:"go1.13.11"}`
func main()   {
   s := `{"Version":"v3.0", "GitCommit":"", "GitTreeState":"", "GoVersion":"go1.13.11"}`
   var data map[string]interface{}
   err := json.Unmarshal([]byte(s), &data)
   if err != nil {
       panic(err)
   }

   fmt.Println("version =", data["Version"])
   //innermap, ok := data["version.BuildInfo"].(map[string]interface{})
   //if !ok {
   //    panic("inner map is not a map!")
   //}

   //fmt.Println("innermap.foo =", innermap["Version"])
}
