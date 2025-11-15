package pengambilandata

type NotificationService interface {
	PushNotification(email, message string) error
}
