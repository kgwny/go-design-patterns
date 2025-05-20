package main

import "fmt"

// Notification は通知の共通インターフェース
type Notification interface {
	Send(message string)
}

// EmailNotification は Email を使った通知の実装
type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Sending Email with message:", message)
}

// SMSNotification は SMS を使った通知の実装
type SMSNotification struct{}

func (s SMSNotification) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

// NotificationFactory は通知タイプに応じて適切な Notification を生成するファクトリーメソッド
func NotificationFactory(notificationType string) Notification {
	switch notificationType {
	case "email":
		return EmailNotification{}
	case "sms":
		return SMSNotification{}
	default:
		panic("Unknown notification type:" + notificationType)
	}
}

// main 関数はファクトリーメソッドを使って通知を送信する
func main() {
	email := NotificationFactory("email")
	email.Send("Welcome to Go Factory Method!")

	sms := NotificationFactory("sms")
	sms.Send("Your code has been verified.")
}
