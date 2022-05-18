package test

import (
	"promo-tables/utils/mock"
	"testing"
)

func TestMockUUID(t *testing.T) {
	if !ValidateUUID(mock.GetUUID()) {
		t.Errorf("mock.GetUUID() returns %v, not a valid uuid", mock.GetUUID())
	}
}

func TestEmail(t *testing.T) {
	name := mock.GetName()
	nickname := mock.GetNickname(name)
	email := mock.GetEmail(nickname)
	if !ValidateEmail(email) {
		t.Errorf("mock.GetEmail() returns %v, not a valid email", email)
	}
}

func TestPassword(t *testing.T) {
	password := mock.GetPassword(16)
	if len(password) != 16 {
		t.Errorf("mock.GetPassword() returns %v, not a valid password", password)
	}
	hashedPassword := mock.GetSHA256Password(password)
	if len(hashedPassword) != 64 {
		t.Errorf("mock.GetSHA256Password() returns %v, not a valid password", hashedPassword)
	}
}

func TestMockTables(t *testing.T) {
	// need sequence
	handler := GetServiceHandler()
	InitMockTable(handler)
	if !CheckTableExists(handler) {
		t.Errorf("[FAIL] InitMockTable()")
	}
	t.Log("[PASS] InitMockTable()")

	CheckUserTable(handler, t)
	t.Log("[PASS] CheckUserTable()")

	CheckCategoryTable(handler, t)
	t.Log("[PASS] CheckCategoryTable()")

	CheckPromoCategoryTable(handler, t)
	t.Log("[PASS] CheckPromoCategoryTable()")

	CheckPromoControllerTable(handler, t)
	t.Log("[PASS] CheckPromoControllerTable()")

	CheckProductTable(handler, t)
	t.Log("[PASS] CheckProductTable()")
}

// // Looks good, but Not Work as test function
// type Handler services.Handler
// func (h *Handler) TestUserTable(t *testing.T) {
// 	user_count := 10
// 	h.service.AddUsers(user_count)
// 	users := h.service.GetUsers()
// 	if len(users) != user_count {
// 		t.Errorf("user table has %v users, not %v", len(users), user_count)
// 	}

// 	for _, user := range users {
// 		if !ValidateUUID(user.UUID) {
// 			t.Errorf("user.UUID is not a valid uuid")
// 		}
// 		if !ValidateEmail(user.Email) {
// 			t.Errorf("user.Email is not a valid email")
// 		}
// 		if !ValidateSHA256Password(user.Password) {
// 			t.Errorf("user.Password is not a valid password")
// 		}
// 	}
// }
