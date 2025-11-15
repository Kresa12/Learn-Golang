package pengambilandata

import "fmt"

type EmailNotifier struct{}

func (e *EmailNotifier) PushNotification(email, message string) error {
	fmt.Printf("Mengirim email ke %s: %s\n", email, message)
	return nil
}
