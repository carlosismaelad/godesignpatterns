package main

import (
	"carlosismaelad.com/godesignpatterns/abstractfactory/notification"
)

func main() {
	var factory notification.NotificationFactory

	factory = notification.EmailFactory{From: "noreply@example.com"}
	notifier := factory.CreateNotifier()
	logger := factory.CreateLogger()
	notifier.Send("user@example.com", "Welcome via Email!")
	logger.Log("Email sent successfully.")

	factory = notification.SMSFactory{Number: "+5511999999999"}
	notifier = factory.CreateNotifier()
	logger = factory.CreateLogger()
	notifier.Send("+5511888888888", "Welcome via SMS!")
	logger.Log("SMS sent successfully.")
}