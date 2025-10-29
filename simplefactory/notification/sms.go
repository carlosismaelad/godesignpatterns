package notification

import "fmt"

type SMSNotifier struct {
	SenderNumber string
}

func (s SMSNotifier) Send(to, message string) error {
	fmt.Printf("📱 Enviando SMS de %s para %s: %s\n", s.SenderNumber, to, message)
	return nil
}