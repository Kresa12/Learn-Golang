package latihan

type UserDataStore interface {
	GetUserName(userID string) (string, error)
}
