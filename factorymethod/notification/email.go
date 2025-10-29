package notification

import "fmt"

type EmailNotifier struct {
    From string
}

func (e EmailNotifier) Send(to, message string) error {
    fmt.Printf("📧 Sending email from %s to %s: %s\n", e.From, to, message)
    return nil
}