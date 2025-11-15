package latihan

type NotificationService interface {
	PushNotification(email, message string) error
}
