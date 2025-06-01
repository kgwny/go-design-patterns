package main

import "fmt"

// Flyweight (共有されるオブジェクト)
type TreeType struct {
	name    string
	color   string
	texture string
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf(
		"Drawing tree '%s' of color '%s' with texture '%s' at (%d, %d)\n",
		t.name, t.color, t.texture, x, y)
}

// Flyweight Factory (TreeType を再利用するためのファクトリー)
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{treeTypes: make(map[string]*TreeType)}
}

func (f *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "_" + color + "_" + texture
	if tree, exists := f.treeTypes[key]; exists {
		return tree
	}
	tree := &TreeType{name: name, color: color, texture: texture}
	f.treeTypes[key] = tree
	return tree
}

// Context (座標などの外部状態を持つオブジェクト)
type Tree struct {
	x, y     int
	treeType *TreeType
}

func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{x: x, y: y, treeType: treeType}
}

func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}

// Forest (ツリーを管理するクライアント)
type Forest struct {
	trees   []*Tree
	factory *TreeFactory
}

func NewForest(factory *TreeFactory) *Forest {
	return &Forest{trees: []*Tree{}, factory: factory}
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.factory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}

func main() {
	factory := NewTreeFactory()
	forest := NewForest(factory)

	// 同じ TreeType を使って複数の Tree を作成
	forest.PlantTree(10, 20, "Oak", "Green", "Rough")
	forest.PlantTree(15, 25, "Oak", "Green", "Rough")
	forest.PlantTree(30, 40, "Pine", "Dark Green", "Smooth")

	forest.Draw()

	// 登録された TreeType の数を表示 (Flyweight による共有を確認)
	fmt.Printf("Total TreeTypes created: %d\n", len(factory.treeTypes))
}
