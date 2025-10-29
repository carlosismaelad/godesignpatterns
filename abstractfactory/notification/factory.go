package notification

import "fmt"

type NotificationFactory interface {
	CreateNotifier() Notifier
	CreateLogger() Logger
}

type Logger interface {
	Log(message string)
}

type EmailLogger struct{}

func (l EmailLogger) Log(message string) {
	fmt.Println("[Email Log]:", message)
}

type SMSLogger struct{}

func (l SMSLogger) Log(message string) {
	fmt.Println("[SMS Log]:", message)
}

type EmailFactory struct {
	From string
}

func (f EmailFactory) CreateNotifier() Notifier {
	return EmailNotifier(f)
}

func (f EmailFactory) CreateLogger() Logger {
	return EmailLogger{}
}

type SMSFactory struct {
	Number string
}

func (f SMSFactory) CreateNotifier() Notifier {
	return SMSNotifier(f)
}

func (f SMSFactory) CreateLogger() Logger {
	return SMSLogger{}
}