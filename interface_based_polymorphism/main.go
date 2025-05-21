package main

import (
	"fmt"
)

// Go における interface を使った動的ディスパッチ（Duck Typing 的発想） を示すシンプルなサンプルプログラム
// ある振る舞い（メソッド）を持つ限り、それを使う側は具体的な型を気にせず扱えることを表現している

// Speaker インターフェースは Speak メソッドを持つ型を要求する
type Speaker interface {
	Speak() string
}

// Dog 構造体
type Dog struct {
	Name string
}

// Dog は Speak メソッドを持つ -> Speaker インターフェースを満たす
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// Cat 構造体
type Cat struct {
	Name string
}

// Cat も Speak メソッドを持つ -> Speaker インターフェースを満たす
func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

// Robot 構造体
type Robot struct {
	Model string
}

// Robot も Speak メソッドを持つ -> Speaker インターフェースを満たす
func (r Robot) Speak() string {
	return "Robot " + r.Model + ": Beep Boop"
}

// introduce 関数は Speaker インターフェースを引数に取る
func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Pochi"}
	cat := Cat{Name: "Tama"}
	robot := Robot{Model: "XJ-9"}

	// 具体的な型に関係なく、Speaker インターフェースを満たしていれば使える
	introduce(dog)
	introduce(cat)
	introduce(robot)
}
