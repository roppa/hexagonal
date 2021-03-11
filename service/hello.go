package service

import (
	hex "github.com/roppa/hexaganol"
)

type hello struct{}

// NewService creates a new hello service
func NewService() *hello {
	return &hello{}
}

// SayHello generates a hello message
func (h *hello) SayHello(args hex.HelloArgs) hex.HelloResponse {
	return hex.HelloResponse{Message: "Hello " + args.Name}
}
