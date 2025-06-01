package main

import "fmt"

// Prototype インターフェース
type Prototype interface {
	Clone() Prototype
	GetInfo() string
}

// ConcretePtototypeA は prototype を実装する構造体のひとつ
type ConcretePrototypeA struct {
	Name  string
	Value int
}

func (p *ConcretePrototypeA) Clone() Prototype {
	// 値のコピーを作成して返す
	clone := *p
	return &clone
}

func (p *ConcretePrototypeA) GetInfo() string {
	return fmt.Sprintf("PrototypeA: Name=%s, Value=%d", p.Name, p.Value)
}

// ConcretePrototypeB も prototype を実装する構造体となる
type ConcretePrototypeB struct {
	Category string
	Amount   float64
}

func (p *ConcretePrototypeB) Clone() Prototype {
	clone := *p
	return &clone
}

func (p *ConcretePrototypeB) GetInfo() string {
	return fmt.Sprintf("PrototypeB: Category=%s, Amount=%.2f", p.Category, p.Amount)
}

func main() {
	// オリジナルのインスタンスを作成
	protoA := &ConcretePrototypeA{Name: "OriginalA", Value: 100}
	protoB := &ConcretePrototypeB{Category: "Finance", Amount: 99.99}

	// 複製を作成
	cloneA := protoA.Clone()
	cloneB := protoB.Clone()

	// オリジナルとクローンの情報表示
	fmt.Println("Original A:", protoA.GetInfo())
	fmt.Println("Cloned A:", cloneA.GetInfo())

	fmt.Println("Original B:", protoB.GetInfo())
	fmt.Println("Cloned B:", cloneB.GetInfo())

	// クローンを変更してオリジナルに影響がないことを確認
	cloneA.(*ConcretePrototypeA).Value = 999
	cloneB.(*ConcretePrototypeB).Amount = 0.01

	fmt.Println("\nAfter modification:")
	fmt.Println("Original A:", protoA.GetInfo())
	fmt.Println("Cloned A:", cloneA.GetInfo())

	fmt.Println("Original B:", protoB.GetInfo())
	fmt.Println("Cloned B:", cloneB.GetInfo())
}
