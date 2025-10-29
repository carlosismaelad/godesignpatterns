package notification

import "fmt"

type SMSNotifier struct {
	Number string
}

func (s SMSNotifier) Send(to, message string) error {
	fmt.Printf("📱 Sending SMS from %s to %s: %s\n", s.Number, to, message)
	return nil
}