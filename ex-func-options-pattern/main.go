package main

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type OptsFunc func(*Opts)

func DefaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

type server struct {
	Opts
}

func NewServer(opts ...OptsFunc) *server {
	o := DefaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &server{
		Opts: o,
	}
}

func withMaxConn(n int) OptsFunc {
	return func(o *Opts) {
		o.maxConn = n
	}
}

func withTLS(o *Opts) {
	o.tls = true
}

func main() {
	server := NewServer(withTLS, withMaxConn(100))
	fmt.Printf("%+v\n", server)
}
