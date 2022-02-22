package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/changpil/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func startWebservice(port, numberOfretries int) (int, error) {
	fmt.Println("Starting server ...")

	fmt.Println("Server started on port", port, numberOfretries)
	// return port, nil
	return port, errors.New("Something went wrong")
}
