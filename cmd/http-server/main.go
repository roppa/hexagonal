package main

import (
	"github.com/labstack/echo"
	shttp "github.com/roppa/hexaganol/transports/http"
)

func main() {
	server := echo.New()
	shttp.NewHelloHandler().Routes(server.Group(""))
	server.Start(":3001")
}
