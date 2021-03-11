package service

import "fmt"

// NewMessage generates a hello message
func NewMessage(args HelloArgs) HelloResponse {
	fmt.Println("= NewMessage ==============================================================")
	fmt.Println("===============================================================")
	fmt.Println(args)
	fmt.Println("===============================================================")
	fmt.Println("===============================================================")
	return HelloResponse{Message: "Hello " + args.Name}
}
