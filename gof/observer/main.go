package main

import "fmt"

// Observer インターフェース
type Observer interface {
	Update(message string)
}

// Subject インターフェース
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll(message string)
}

// ConreteSubject は Subject の具体的な実装
type ConcreteSubject struct {
	observers []Observer
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Unregister(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			// スライスから削除
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyAll(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserverA
type ConcreteObserverA struct {
	ID string
}

func (o *ConcreteObserverA) Update(message string) {
	fmt.Printf("[Observer A - %s] 受信: %s\n", o.ID, message)
}

// ConcreteObserverB
type ConcreteObserverB struct {
	ID string
}

func (o *ConcreteObserverB) Update(message string) {
	fmt.Printf("[Observer B - %s] 受信: %s\n", o.ID, message)
}

func main() {
	subject := NewConcreteSubject()

	observer1 := &ConcreteObserverA{ID: "001"}
	observer2 := &ConcreteObserverB{ID: "XYZ"}

	subject.Register(observer1)
	subject.Register(observer2)

	fmt.Println("-- 通知 1回目 --")
	subject.NotifyAll("イベント発生！")

	subject.Unregister(observer1)

	fmt.Println("-- 通知 2回目 (observer1 を削除後) --")
	subject.NotifyAll("別のイベント発生！")
}
