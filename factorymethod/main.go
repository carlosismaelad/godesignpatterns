package main

import (
	"carlosismaelad.com/godesignpatterns/factorymethod/notification"
)

func main() {
    // Usando a factory de e-mail
    var factory notification.NotifierFactory
    factory = notification.EmailFactory{From: "noreply@example.com"}
    notifier := factory.CreateNotifier()
    notifier.Send("user@example.com", "Welcome via Email!")

    // Usando a factory de SMS
    factory = notification.SMSFactory{Number: "+5511999999999"}
    notifier = factory.CreateNotifier()
    notifier.Send("+5511888888888", "Welcome via SMS!")
}