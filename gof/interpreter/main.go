package main

import (
	"fmt"
	"strings"
)

// Expression インターフェース
type Expression interface {
	Interpret(context string) bool
}

// TerminalExpression: 単一のキーワードに対応
type TerminalExpression struct {
	data string
}

func (t *TerminalExpression) Interpret(context string) bool {
	return strings.Contains(context, t.data)
}

// OrExpression: 論理 OR
type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context)
}

// AndExpression: 論理 AND
type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context)
}

// クライアント用の構成関数
func getMaleExpression() Expression {
	robert := &TerminalExpression{data: "Robert"}
	john := &TerminalExpression{data: "John"}
	return &OrExpression{expr1: robert, expr2: john}
}

func getMarriedWomanExpression() Expression {
	julie := &TerminalExpression{data: "Julie"}
	married := &TerminalExpression{data: "Married"}
	return &AndExpression{expr1: julie, expr2: married}
}

func main() {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()

	fmt.Println("John is male? ->", isMale.Interpret("John"))
	fmt.Println("Julie is a married woman? ->", isMarriedWoman.Interpret("Married Julie"))
	fmt.Println("Julie (not married) is a married woman? ->", isMarriedWoman.Interpret("Julie"))
}
