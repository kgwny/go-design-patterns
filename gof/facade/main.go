package main

import "fmt"

// サブシステム1: DVDプレイヤー
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player: ON")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("DVD Player: Playing \"%s\"\n", movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player: OFF")
}

// サブシステム2: プロジェクター
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector: ON")
}

func (p *Projector) WideScreenMode() {
	fmt.Println("Projector: Set to widescreen mode")
}

func (p *Projector) Off() {
	fmt.Println("Projector: OFF")
}

// サブシステム3: スピーカー
type Speaker struct{}

func (s *Speaker) On() {
	fmt.Println("Speaker: ON")
}

func (s *Speaker) SetVolume(level int) {
	fmt.Printf("Speaker: Volume set to %d\n", level)
}

func (s *Speaker) Off() {
	fmt.Println("Speaker: OFF")
}

// Facade: ホームシアター全体の走査を簡素化する
type HomeTheaterFacade struct {
	dvd       *DVDPlayer
	projector *Projector
	speaker   *Speaker
}

func NewHomeTheaterFacade(dvd *DVDPlayer, projector *Projector, speaker *Speaker) *HomeTheaterFacade {
	return &HomeTheaterFacade{dvd, projector, speaker}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.speaker.On()
	h.speaker.SetVolume(5)
	h.projector.On()
	h.projector.WideScreenMode()
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.dvd.Off()
	h.projector.Off()
	h.speaker.Off()
}

func main() {
	// 各コンポーネントを生成
	dvd := &DVDPlayer{}
	projector := &Projector{}
	speaker := &Speaker{}

	// Facade を生成
	homeTheater := NewHomeTheaterFacade(dvd, projector, speaker)

	// 映画鑑賞開始と終了
	homeTheater.WatchMovie("The Matrix")
	fmt.Println()
	homeTheater.EndMovie()
}
