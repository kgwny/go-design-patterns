package main

import "fmt"

// Handler インターフェース
// 各ハンドラーは責務を持ち、処理できない場合は次のハンドラーに渡する
// リクエストの種類に応じて異なるハンドラーが対応する
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler: 共通の構造体 (次のハンドラーを保持)
type BaseHandler struct {
	next Handler
}

// SetNext: 次のハンドラーを設定
func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

// CallNext: 次のハンドラーへ処理を渡す
func (b *BaseHandler) CallNext(request string) {
	if b.next != nil {
		b.next.Handle(request)
	} else {
		fmt.Println("No handler could process the request:", request)
	}
}

// ConcreteHandlerA: 特定のリクエストを処理
type HandlerA struct {
	BaseHandler
}

func (h *HandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("HandlerA processed request:", request)
	} else {
		h.CallNext(request)
	}
}

// ConcreteHandlerB: 別のリクエストを処理
type HandlerB struct {
	BaseHandler
}

func (h *HandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("HandlerB processed request:", request)
	} else {
		h.CallNext(request)
	}
}

// ConcreteHandlerC: デフォルトのハンドラー
type HandlerC struct {
	BaseHandler
}

func (h *HandlerC) Handle(request string) {
	if request == "C" {
		fmt.Println("HandlerC processed request:", request)
	} else {
		h.CallNext(request)
	}
}

func main() {
	// 各ハンドラーを作成
	handlerA := &HandlerA{}
	handlerB := &HandlerB{}
	handlerC := &HandlerC{}

	// チェーンを構築
	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	// 各種リクエストを処理
	requests := []string{"A", "B", "C", "D"}
	for _, req := range requests {
		fmt.Println("Sending request:", req)
		handlerA.Handle(req)
		fmt.Println()
	}
}
