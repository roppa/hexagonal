package hexaganol

type (
	// HelloArgs is the parameter structure for saying hello.
	HelloArgs struct {
		Name string
	}

	// HelloResponse is the greeting
	HelloResponse struct {
		Message string
	}
)

// HelloService when implemented will contain all business logic.
type HelloService interface {
	SayHello(args HelloArgs) HelloResponse
}
