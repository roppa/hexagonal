package http

import (
	"net/http"

	"github.com/labstack/echo"
	hex "github.com/roppa/hexaganol"
)

type helloHandler struct {
	svc hex.HelloService
}

// NewHelloHandler returns hello handler.
func NewHelloHandler(svc hex.HelloService) *helloHandler {
	return &helloHandler{svc: svc}
}

func (h *helloHandler) Routes(g *echo.Group) {
	g.GET("/hello", h.sayHello)
}

func (h *helloHandler) sayHello(e echo.Context) error {
	args := hex.HelloArgs{Name: e.QueryParam("name")}
	return e.JSON(http.StatusOK, h.svc.SayHello(args))
}
