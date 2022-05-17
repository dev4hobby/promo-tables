package services

import (
	"promo-tables/database/models"
	"promo-tables/utils/mock"
)

func (h *Handler) AddUsers(howMany int) error {
	return h.db.AddUsers(GetUser(howMany))
}

func GetUser(count int) []models.User {
	users := []models.User{}

	for i := 0; i < count; i++ {
		name := mock.GetName()
		nickname := mock.GetNickname(name)
		email := mock.GetEmail(nickname)
		users = append(users, models.User{
			UUID:     mock.GetUUID(),
			Name:     name,
			Nickname: nickname,
			Email:    email,
			Password: mock.GetSHA256Password(mock.GetPassword(16)),
			Sex:      mock.GetSex(),
			Birth:    mock.GetBirth(),
			Phone:    mock.GetPhone(),
		})
	}
	return users
}
