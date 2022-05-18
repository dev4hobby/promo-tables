package test

import "regexp"

func ValidateUUID(uuid string) bool {
	match, _ := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", uuid)
	return len(uuid) == 36 && match
}

func ValidateEmail(email string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return match
}

func ValidateSHA256Password(password string) bool {
	match, _ := regexp.MatchString("^[a-f0-9]{64}$", password)
	return match
}
