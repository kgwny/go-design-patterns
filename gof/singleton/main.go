package main

import (
	"fmt"
	"sync"
)

// Singleton 構造体の定義
type Singleton struct {
	Value string
}

var instance *Singleton
var once sync.Once

// GetInstance は Singleton のインスタンスを返します (初回のみ初期化される)
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance...")
		instance = &Singleton{Value: "I am the only one!"}
	})
	return instance
}

func main() {
	// 2つのインスタンスを取得 (実際には同じもの)
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1.Value)
	fmt.Println("s1 == s2?", s1 == s2) // true が出力される
}
