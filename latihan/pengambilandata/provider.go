package pengambilandata

type UserDataStore interface {
	GetUserName(userID string) (string, error)
}
