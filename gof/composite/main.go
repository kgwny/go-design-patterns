package main

import "fmt"

// Component インターフェース (共通のインターフェース)
type Component interface {
	Name() string
	Display(indent int)
}

// File
type File struct {
	name string
}

// Directory (コンポジット)
type Directory struct {
	name       string
	components []Component
}

// インデント用のヘルパー関数
func getIndent(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Display(indent int) {
	fmt.Printf("%s- File: %s\n", getIndent(indent), f.name)
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Add(c Component) {
	d.components = append(d.components, c)
}

func (d *Directory) Display(indent int) {
	fmt.Printf("%s+ Dicrectory: %s\n", getIndent(indent), d.name)
	for _, c := range d.components {
		c.Display(indent + 2)
	}
}

// メイン関数
func main() {
	// ファイルとディレクトリを作成する
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	file3 := &File{name: "file3.log"}

	dir1 := &Directory{name: "dir1"}
	dir2 := &Directory{name: "dir2"}

	// 構造を構築する
	dir1.Add(file1)
	dir1.Add(file2)

	dir2.Add(file3)
	dir2.Add(dir1) // dir2 のなかに dir1 を追加 (ネスト)

	// 全体を表示する
	dir2.Display(0)
}
