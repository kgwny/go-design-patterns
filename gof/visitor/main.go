package main

import (
	"fmt"
	"math"
)

// Visitor
type ShapeVisitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

// Element
type Shape interface {
	Accept(visitor ShapeVisitor)
}

// Concrete Elements
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

// Concrete Visitors
type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("Circle Area: %.2f\n", area)
}

func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("Rectangle Area: %.2f\n", area)
}

type SVGExporter struct{}

func (s *SVGExporter) VisitCircle(c *Circle) {
	fmt.Printf("Exporting Circle as SVG: <circle r=\"%.2f\" />\n", c.Radius)
}

func (s *SVGExporter) VisitRectangle(r *Rectangle) {
	fmt.Printf("Exporting Rectangle as SVG: <rect width=\"%.2f\" height=\"%.2f\" />\n", r.Width, r.Height)
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Width: 4, Height: 6},
	}

	fmt.Println("--- Calculating Area ---")
	areaCalculator := &AreaCalculator{}
	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}

	fmt.Println("\n--- Exporting to SVG ---")
	svgExporter := &SVGExporter{}
	for _, shape := range shapes {
		shape.Accept(svgExporter)
	}
}
