package latihan

import "fmt"

type WhatsAppNotifier struct{}

func (w *WhatsAppNotifier) PushNotification(email, message string) error {
	fmt.Printf("Mengirim WhatsApp ke %s: %s\n", email, message)
	return nil
}
