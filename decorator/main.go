package main

import "fmt"

// Component インターフェース
type Coffee interface {
	GetCost() int
	GetDescription() string
}

// ConcreteComponent
type SimpleCoffee struct{}

func (s *SimpleCoffee) GetCost() int {
	return 300
}

func (s *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

// Decorator 抽象構造体 (埋め込み)
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) GetCost() int {
	return d.coffee.GetCost()
}

func (d *CoffeeDecorator) GetDescription() string {
	return d.coffee.GetDescription()
}

// MilkDecorator (ConcreteDecorator)
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(c Coffee) *MilkDecorator {
	return &MilkDecorator{CoffeeDecorator{c}}
}

func (m *MilkDecorator) GetCost() int {
	return m.coffee.GetCost() + 50
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

// SugarDecorator (ConcreteDecorator)
type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(c Coffee) *SugarDecorator {
	return &SugarDecorator{CoffeeDecorator{c}}
}

func (s *SugarDecorator) GetCost() int {
	return s.coffee.GetCost() + 30
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + ", Sugar"
}

func main() {
	// シンプルなコーヒー
	var coffee Coffee = &SimpleCoffee{}
	fmt.Println(coffee.GetDescription(), ":", coffee.GetCost(), "yen")

	// ミルクを追加
	coffee = NewMilkDecorator(coffee)
	fmt.Println(coffee.GetDescription(), ":", coffee.GetCost(), "yen")

	// さらに砂糖を追加
	coffee = NewSugarDecorator(coffee)
	fmt.Println(coffee.GetDescription(), ":", coffee.GetCost(), "yen")
}
