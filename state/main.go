package main

import "fmt"

// State インターフェース：全ての状態が実装する共通インターフェース
type State interface {
	Publish(d *Document)
	Name() string
}

// Draft 状態
type Draft struct{}

func (s *Draft) Publish(d *Document) {
	fmt.Println("Transitioning from Draft to Review...")
	d.SetState(&Review{})
}

func (s *Draft) Name() string {
	return "Draft"
}

// Review 状態
type Review struct{}

func (s *Review) Publish(d *Document) {
	fmt.Println("Transitioning from Review to Published...")
	d.SetState(&Published{})
}

func (s *Review) Name() string {
	return "Review"
}

// Published 状態
type Published struct{}

func (s *Published) Publish(d *Document) {
	fmt.Println("Already Published. No further transitions.")
}

func (s *Published) Name() string {
	return "Published"
}

// Document コンテキスト構造体：状態を保持し、状態に移譲する
type Document struct {
	state State
}

func NewDocument() *Document {
	return &Document{state: &Draft{}}
}

func (d *Document) Publish() {
	d.state.Publish(d)
}

func (d *Document) SetState(s State) {
	d.state = s
}

func (d *Document) StateName() string {
	return d.state.Name()
}

func main() {
	doc := NewDocument()
	fmt.Println("Current State:", doc.StateName())

	doc.Publish()
	fmt.Println("Current State:", doc.StateName())

	doc.Publish()
	fmt.Println("Current State:", doc.StateName())

	doc.Publish() // Published状態なので変化なし
	fmt.Println("Current State:", doc.StateName())
}
