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
func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Routes(g *echo.Group) {
	g.GET("/hello", h.sayHello)
	g.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
}

func (h *helloHandler) sayHello(e echo.Context) error {
	// args := hex.HelloArgs{Name: e.Param("name")}
	return e.JSON(http.StatusOK, "hello "+e.QueryParam("name"))
}
