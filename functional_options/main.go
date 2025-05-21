package main

import "fmt"

// Functional Options パターン（構成オプションを関数で渡す）
// 構造体の初期化時に設定オプションを柔軟に渡したいとき、Go では「関数を使って構成する」ことが多い

type Server struct {
	Host string
	Port int
}

type Option func(*Server)

func WithHost(h string) Option {
	return func(s *Server) {
		s.Host = h
	}
}

func WithPort(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}

func NewServer(opts ...Option) *Server {
	s := &Server{Host: "localhost", Port: 8080}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	s := NewServer(WithHost("example.com"), WithPort(9000))
	fmt.Println(s)
}
