package main

import (
	"fmt"
	"mongodbinstallation/router"
	"net/http"
)

func main() {
fmt.Println("API Project")
r:=router.Router()
http.ListenAndServe(":4000",r)
}

