package latihan

import "fmt"

type ApiStore struct{}

func (a *ApiStore) GetUserName(userID string) (string, error) {
	return fmt.Sprintf("Mengambil %s dari API Store \n", userID), nil
}
