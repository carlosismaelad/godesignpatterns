package notification

type NotifierFactory interface {
    CreateNotifier() Notifier
}

type EmailFactory struct {
    From string
}

func (f EmailFactory) CreateNotifier() Notifier {
    return EmailNotifier(f)
}

type SMSFactory struct {
    Number string
}

func (f SMSFactory) CreateNotifier() Notifier {
    return SMSNotifier(f)
}