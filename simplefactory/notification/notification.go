package notification

type Notifier interface {
	Send(to, message string) error
}

func SendNotification(n Notifier, to, message string) error {
	return n.Send(to, message)
}