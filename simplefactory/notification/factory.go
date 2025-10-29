package notification

import (
	"fmt"
)

func NewNotifier(kind string) (Notifier, error) {
	switch kind {
	case "email":
		return EmailNotifier{From: "noreply@exemplo.com"}, nil
	case "sms":
		return SMSNotifier{SenderNumber: "+55 11 99999-9999"}, nil
	default:
		return nil, fmt.Errorf("tipo de notificador desconhecido: %s", kind)
	}
}