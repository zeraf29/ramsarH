package main

import (
	"fmt"
	_ "github.com/labstack/echo/middleware"
	"github.com/zeraf29/ramsarH/router"
)

func main() {
	//debug mode
	debug := true

	//Declare Router
	eRouter := router.Router()

	fmt.Print("------Starting RamsarH Project......-----")

	if debug {
		// Normal http server
		eRouter.Logger.Fatal(eRouter.Start(":9091"))
	} else {
		// Security https Server
		//eRouter.Logger.Fatal(eRouter.StartTLS(":9091", "cert.pem", "privkey.pem"))
	}
}
