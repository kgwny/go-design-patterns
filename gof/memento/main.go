package main

import "fmt"

// Memento は Editor の状態を保存する構造体
type Memento struct {
	content string
}

// Editor はテキストを編集する構造体
type Editor struct {
	content string
}

// CreateMemento は現在の状態を Memento に保存する
func (e *Editor) CreateMemento() *Memento {
	return &Memento{content: e.content}
}

// Restore は Memento から状態を復元する
func (e *Editor) Restore(m *Memento) {
	e.content = m.content
}

// TypeText はテキストを追加する
func (e *Editor) TypeText(text string) {
	e.content += text
}

// Content は現在のテキストを返す
func (e *Editor) Content() string {
	return e.content
}

// History は Memento のスタックを管理する構造体
type History struct {
	mementos []*Memento
}

// Push は Memento をスタックに追加する
func (h *History) Push(m *Memento) {
	h.mementos = append(h.mementos, m)
}

// Pop は最後に追加された Memento を取り出す
func (h *History) Pop() *Memento {
	if len(h.mementos) == 0 {
		return nil
	}
	last := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return last
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.TypeText("Hello")
	history.Push(editor.CreateMemento())

	editor.TypeText(", world!")
	history.Push(editor.CreateMemento())

	editor.TypeText(" Goodbye!")
	fmt.Println("Current Content:", editor.Content())

	// Undo once
	editor.Restore(history.Pop())
	fmt.Println("After Undo 1:", editor.Content())

	// Undo again
	editor.Restore(history.Pop())
	fmt.Println("After Undo 2:", editor.Content())
}
