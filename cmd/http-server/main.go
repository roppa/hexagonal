package main

import (
	"github.com/labstack/echo"
	"github.com/roppa/hexaganol/service"
	shttp "github.com/roppa/hexaganol/transports/http"
)

func main() {
	server := echo.New()
	shttp.NewHelloHandler(service.NewService()).Routes(server.Group(""))
	server.Start(":3001")
}
