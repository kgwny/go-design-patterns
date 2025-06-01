package main

import "fmt"

// クライアントが期待しているインターフェース (日本のコンセント)
type JapaneseSocket interface {
	InsertIntoSocket()
}

// 適合させたい既存のクラス (ヨーロッパプラグ)
type EuropeanPlug struct{}

func (e *EuropeanPlug) InsertIntoEuropeanSocket() {
	fmt.Println("ヨーロッパのコンセントに接続されました")
}

// Adapter:JapaneseSocket インターフェースに適合するようにする
type PlugAdapter struct {
	europeanDevice *EuropeanPlug
}

func (a *PlugAdapter) InsertIntoSocket() {
	fmt.Println("アダプターを使って変換中...")
	a.europeanDevice.InsertIntoEuropeanSocket()
}

// クライアントコード
func ConnectToJapaneseSocket(socket JapaneseSocket) {
	socket.InsertIntoSocket()
}

func main() {
	// ヨーロッパのデバイスを作成
	euroDevice := &EuropeanPlug{}

	// アダプターを作成して日本のソケットに接続できるようにする
	adapter := &PlugAdapter{europeanDevice: euroDevice}

	// ヨーロッパのプラグを日本のコンセントに接続 (アダプター経由)
	ConnectToJapaneseSocket(adapter)
}
