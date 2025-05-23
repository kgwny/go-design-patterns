package main

import "fmt"

// === 抽象製品インターフェース ===

type Button interface {
	Render()
}

type Checkbox interface {
	Check()
}

// === 抽象ファクトリ ===

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// === Mac の実装 ===

type MacButton struct{}

func (b *MacButton) Render() {
	fmt.Println("Render a Mac style button")
}

type MacCheckbox struct{}

func (c *MacCheckbox) Check() {
	fmt.Println("CHeck a Mac style checkbox")
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// === Windows の実装 ===

type WinButton struct{}

func (b *WinButton) Render() {
	fmt.Println("Render a Windows style button")
}

type WinCheckbox struct{}

func (c *WinCheckbox) Check() {
	fmt.Println("Check a Windows style checkbox")
}

type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (f *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

// === クライアントコード ===

func main() {
	var factory GUIFactory

	os := "mac" // "windows" に切り替えると WinFactory を使う

	if os == "mac" {
		factory = &MacFactory{}
	} else {
		factory = &WinFactory{}
	}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Render()
	checkbox.Check()
}
