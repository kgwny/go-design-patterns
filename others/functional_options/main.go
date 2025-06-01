package main

import "fmt"

// Functional Options パターン
// 構造体の初期化時に、設定用の関数（Option関数）を受け取る仕組みにすることにより、構造体のフィールドを柔軟に設定できる。

// Goにはコンストラクタのオーバーロードがありません。以下のような課題を解決できます：
// オプションが多いとコンストラクタの引数が煩雑になる
// 変更や拡張がしにくい
// 呼び出し側の可読性が下がる

// Functional Options パターンを使うことで以下のようなメリットがあります：
// 柔軟性：必要なオプションだけを設定できる
// 拡張性：後から新しいオプションを追加しやすい
// 可読性：呼び出し側のコードが明示的で読みやすい

type Server struct {
	Host string
	Port int
}

// 関数を使ってオプションを設定します。
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
