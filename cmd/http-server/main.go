package main

import (
	"github.com/labstack/echo"
	shttp "github.com/roppa/hexaganol/io/http"
	"github.com/roppa/hexaganol/service"
)

func main() {
	server := echo.New()
	shttp.NewHelloHandler(service.NewService()).Routes(server.Group(""))
	server.Start(":3001")
}
