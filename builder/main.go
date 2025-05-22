package main

import "fmt"

// Product: 組み立てる対象
type Pizza struct {
	dough   string
	sauce   string
	topping string
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Pizza with %s dough, %s sause, and %s topping.", p.dough, p.sauce, p.topping)
}

// Builder Interface
type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSause() PizzaBuilder
	SetTopping() PizzaBuilder
	Build() *Pizza
}

// ConcreteBuilder
type MargheritaBuilder struct {
	pizza *Pizza
}

func NewMargheritaBuilder() *MargheritaBuilder {
	return &MargheritaBuilder{pizza: &Pizza{}}
}

func (b *MargheritaBuilder) Build() *Pizza {
	return b.pizza
}

func (b *MargheritaBuilder) SetDough() PizzaBuilder {
	b.pizza.dough = "Thin Crust"
	return b
}

func (b *MargheritaBuilder) SetSause() PizzaBuilder {
	b.pizza.sauce = "Tomato Basil"
	return b
}

func (b *MargheritaBuilder) SetTopping() PizzaBuilder {
	b.pizza.topping = "Mozzarella & Basil"
	return b
}

// Director: 組み立ての順序を管理
type Director struct {
	builder PizzaBuilder
}

func (d *Director) SetBuilder(b PizzaBuilder) {
	d.builder = b
}

func (d *Director) Construct() *Pizza {
	return d.builder.SetDough().SetSause().SetTopping().Build()
}

func main() {
	builder := NewMargheritaBuilder()
	director := Director{}
	director.SetBuilder(builder)

	pizza := director.Construct()
	fmt.Println(pizza)
}
