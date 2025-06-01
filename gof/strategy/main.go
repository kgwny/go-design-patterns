package main

import "fmt"

// Strategy interface (戦略インターフェース)
type DiscountStrategy interface {
	ApplyDiscount(price float64) float64
}

// ConcreteStrategy A (通常割引: 10%オフ)
type PercentageDiscount struct{}

func (d *PercentageDiscount) ApplyDiscount(price float64) float64 {
	return price * 0.9 // 10% off
}

// ConcreteStrategy B (定額割引: 500円引き)
type FlatDiscount struct{}

func (d *FlatDiscount) ApplyDiscount(price float64) float64 {
	return price - 500
}

// ConcreteStrategy C (割引なし)
type NoDiscount struct{}

func (d *NoDiscount) ApplyDiscount(price float64) float64 {
	return price
}

// Context (文脈)
// 割引戦略を使用して最終価格を計算する
type PriceCalculator struct {
	strategy DiscountStrategy
}

func (p *PriceCalculator) SetStrategy(strategy DiscountStrategy) {
	p.strategy = strategy
}

func (p *PriceCalculator) Calculate(price float64) float64 {
	if p.strategy == nil {
		return price
	}
	return p.strategy.ApplyDiscount(price)
}

func main() {
	calculator := &PriceCalculator{}

	originalPrice := 3000.0

	// 通常の10%割引
	calculator.SetStrategy(&PercentageDiscount{})
	fmt.Printf("10%% 割引後の価格: %.0f円\n", calculator.Calculate(originalPrice))

	// 定額500円引き
	calculator.SetStrategy(&FlatDiscount{})
	fmt.Printf("500円引き後の価格: %.0f円\n", calculator.Calculate(originalPrice))

	// 割引なし
	calculator.SetStrategy(&NoDiscount{})
	fmt.Printf("割引なし価格: %.0f円\n", calculator.Calculate(originalPrice))
}
