package test

import (
	"promo-tables/database/models"
	"promo-tables/utils/mock"
	"testing"
)

func BenchmarkFunctionA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var user models.User
		CreateNewUserA(&user)
		user.ID = 1
	}
}

func BenchmarkFunctionB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := CreateNewUserB()
		user.ID = 1
	}
}

func CreateNewUserA(user *models.User) {
	name := mock.GetName()
	nickname := mock.GetNickname(name)
	user.Name = name
	user.Nickname = nickname
	user.Email = mock.GetEmail(nickname)
	user.UUID = mock.GetUUID()
	user.Password = mock.GetSHA256Password(mock.GetPassword(16))
	user.Sex = mock.GetSex()
	user.Birth = mock.GetBirth()
}

func CreateNewUserB() models.User {
	name := mock.GetName()
	nickname := mock.GetNickname(name)
	return models.User{
		Name:     name,
		Nickname: nickname,
		Email:    mock.GetEmail(nickname),
		UUID:     mock.GetUUID(),
		Password: mock.GetSHA256Password(mock.GetPassword(16)),
		Sex:      mock.GetSex(),
		Birth:    mock.GetBirth(),
	}
}
