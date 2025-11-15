package latihan

import "errors"

type DatabaseStore struct {
	Data map[string]string
}

func NewDatabaseStore() *DatabaseStore {
	return &DatabaseStore{
		make(map[string]string),
	}
}

func (d *DatabaseStore) GetUserName(userID string) (string, error) {
	if value, ok := d.Data[userID]; ok {
		return value, nil
	}
	return "", errors.New("user tidak ditemukan di dalam database")
}
