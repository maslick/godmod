package godmod

type Hello struct{}

type IHello interface {
	SayHello() string
}

// Say hello.
func (h *Hello) SayHello() string {
	return "Hello World!"
}

// Create new Hello
func New() IHello {
	return &Hello{}
}
