package main

import (
	"fmt"

	"carlosismaelad.com/godesignpatterns/simplefactory/notification"
)

func main() {
	email, err := notification.NewNotifier("email")
	if err != nil {
		fmt.Println(err)
		return
	}

	sms, err := notification.NewNotifier("sms")
	if err != nil {
		fmt.Println(err)
		return
	}

	notification.SendNotification(email, "carlos@example.com", "Bem-vindo ao sistema!")
	notification.SendNotification(sms, "+55 11 91234-5678", "Seu código é 123456.")
}
