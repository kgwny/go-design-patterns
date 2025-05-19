package main

import "fmt"

// Stepper インターフェースはカスタムステップを定義
type Stepper interface {
	Step1()
	Step2()
}

// Template は処理のテンプレートを示す
type Template struct {
	stepper Stepper
}

func (t *Template) Execute() {
	fmt.Println("=== Template Start ===")
	t.stepper.Step1()
	t.CommonStep()
	t.stepper.Step2()
	fmt.Println("=== Template End ===")
}

func (t *Template) CommonStep() {
	fmt.Println("Common Step: Logging or validation")
}

// --- CSVProcessor 実装 ---
type CSVProcessor struct{}

func (c *CSVProcessor) Step1() {
	fmt.Println("CSV Step1: Read CSV file")
}

func (c *CSVProcessor) Step2() {
	fmt.Println("CSV Step2: Parse and store data")
}

// --- JSONPRocessor 実装 ---
type JSONProcessor struct{}

func (j *JSONProcessor) Step1() {
	fmt.Println("JSON Step1: Read JSON data")
}

func (j *JSONProcessor) Step2() {
	fmt.Println("JSON Step2: Convert and save toDB")
}

func main() {
	fmt.Println("Running CSVProcessor:")
	csv := &CSVProcessor{}
	template1 := &Template{stepper: csv}
	template1.Execute()

	fmt.Println()

	fmt.Println("Running JSONProcessor:")
	json := &JSONProcessor{}
	template2 := &Template{stepper: json}
	template2.Execute()
}
