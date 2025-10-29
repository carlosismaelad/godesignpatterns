package notification

type Notifier interface {
	Send(to, message string) error
}