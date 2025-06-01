package main

import "fmt"

// Mediator interface
type ChatRoomMediator interface {
	ShowMessage(user *User, message string)
}

// ConcreteMediator
type ChatRoom struct{}

func (c *ChatRoom) ShowMessage(user *User, message string) {
	fmt.Printf("[%s]: %s\n", user.GetName(), message)
}

// Colleague
type User struct {
	name     string
	mediator ChatRoomMediator
}

func NewUser(name string, mediator ChatRoomMediator) *User {
	return &User{
		name:     name,
		mediator: mediator,
	}
}

func (u *User) Send(message string) {
	u.mediator.ShowMessage(u, message)
}

func (u *User) GetName() string {
	return u.name
}

func main() {
	chatRoom := &ChatRoom{}

	alice := NewUser("Alice", chatRoom)
	bob := NewUser("Bob", chatRoom)

	alice.Send("Hi Bob! How are you?")
	bob.Send("I'm good, thanks Alice!")
}
