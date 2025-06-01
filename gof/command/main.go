package main

import "fmt"

// Command インターフェース
type Command interface {
	Execute()
	Undo()
}

// Receiver：実際の処理を行う構造体
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("ライトをオンにしました。")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("ライトをオフにしました。")
}

// -----------------------------
// Concrete Command：Light を操作
// -----------------------------
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

// -----------------------------
// Invoker：コマンドを実行する構造体
// -----------------------------
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

// Main
func main() {
	// レシーバー
	light := &Light{}

	// コマンド
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// インボーカー
	remote := &RemoteControl{}

	// ライトをオンにする
	remote.SetCommand(lightOn)
	remote.PressButton()

	// アンドゥ (ライトをオフにする)
	remote.PressUndo()

	// ライトをオフにする
	remote.SetCommand(lightOff)
	remote.PressButton()

	// アンドゥ (ライトをオンにする)
	remote.PressUndo()
}
