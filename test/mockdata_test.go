package test

import (
	"promo-tables/utils"
	"promo-tables/utils/mock"
	"regexp"
	"testing"
)

func TestMockName(t *testing.T) {
	name := mock.GetName()
	if !utils.Contains(mock.Names, name) {
		t.Errorf("mock.GetName() returns %v, not not included name list", name)
	}
}

func TestMockUUID(t *testing.T) {
	uuid := mock.GetUUID()
	match, _ := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", uuid)
	if len(uuid) != 36 || !match {
		t.Errorf("mock.GetUUID() returns %v, not a valid UUID", uuid)
	}
}

func TestEmail(t *testing.T) {
	name := mock.GetName()
	nickname := mock.GetNickname(name)
	email := mock.GetEmail(nickname)
	//go email regexp
	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if !match {
		t.Errorf("mock.GetEmail() returns %v, not a valid email", email)
	}
}
