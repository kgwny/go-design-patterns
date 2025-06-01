package main

import "fmt"

// Image インターフェース：共通のインターフェース
type Image interface {
	Display()
}

// RealImage：実際の画像クラス (重い処理を模倣)
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading image from disk:", filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() {
	fmt.Println("Displaying image:", r.filename)
}

// ImageProxy：Proxy クラス (遅延ロードやアクセス制御を担当)
type ImageProxy struct {
	filename  string
	realImage *RealImage
}

func NewImageProxy(filename string) *ImageProxy {
	return &ImageProxy{filename: filename}
}

func (p *ImageProxy) Display() {
	if p.realImage == nil {
		fmt.Println("Creating RealImage only when needed...")
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func main() {
	fmt.Println("Client: Creating proxy.")
	image := NewImageProxy("sample.jpg")

	fmt.Println("\nClient: First call to Display()")
	image.Display()

	fmt.Println("\nClient: Second call to Display()")
	image.Display()
}
