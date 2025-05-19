package main

import "fmt"

// Book は書籍の情報を表す構造体
type Book struct {
	Title string
}

// BookCollection は Book のスライスを保持する構造体
type BookCollection struct {
	books []Book
}

// AddBook はコレクションに書籍を追加する
func (bc *BookCollection) AddBook(book Book) {
	bc.books = append(bc.books, book)
}

// BookIterator は BookCollection を順番に操作するイテレータとなる
type BookIterator struct {
	books []Book
	index int
}

// CreateIterator は新しい BookIterator を返す
func (bc *BookCollection) CreateIterator() *BookIterator {
	return &BookIterator{
		books: bc.books,
		index: 0,
	}
}

// HasNext は次の要素が存在するかどうかを返す
func (it *BookIterator) HasNext() bool {
	return it.index < len(it.books)
}

// Next は次の Book を返し、インデックスを進める
func (it *BookIterator) Next() *Book {
	if it.HasNext() {
		book := &it.books[it.index]
		it.index++
		return book
	}
	return nil
}

// main 関数：イテレータを使って書籍を順番に表示
func main() {
	collection := &BookCollection{}
	collection.AddBook(Book{Title: "Go言語によるWebアプリ開発"})
	collection.AddBook(Book{Title: "マスタリングGo"})
	collection.AddBook(Book{Title: "Go並行処理実践"})

	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.Title)
	}
}
