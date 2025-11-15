package latihan

import "fmt"

type UserService struct {
	store UserDataStore
}

func NewUserService(store UserDataStore) *UserService {
	return &UserService{
		store,
	}
}

func (s *UserService) ShowUserName(userId string) {
	nama, err := s.store.GetUserName(userId)

	if err != nil {
		fmt.Printf("Gagal mengambil user %s: %v\n", userId, err)
	} else {
		fmt.Printf("Halo, %s!\n", nama)
	}
}
