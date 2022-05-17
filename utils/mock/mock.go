package mock

import (
	"fmt"
	"math"
	"math/rand"
	"promo-tables/utils"
	"time"

	"github.com/google/uuid"
)

var Names = []string{
	"Abel", "Brian", "Catharine", "Dexter",
	"Emily", "Felix", "Gabriel", "Hayden",
	"Iris", "Jake", "Kyle", "Liam", "Mia",
	"Noah", "Owen", "Peter", "Quincy", "Ruby",
	"Sofia", "Tommy", "Ulises", "Vincent",
	"Wade", "Xena", "Yusuf", "Zoe",
}

var MailHost = []string{
	"gmail.com", "naver.com", "kakao.com",
	"hanmail.net", "nate.com", "daum.net",
	"yahoo.com", "outlook.com", "msn.com",
	"hotmail.com", "yahoo.co.kr", "yahoo.co.jp",
}

// KRW based price, if u wanna use it as globally currency,
// you should convert it to others.
var Products = map[string]int{
	"IPhone 13 Pro":  1350000,
	"IPhone 13":      950000,
	"IPhone SE":      590000,
	"IPhone 12":      850000,
	"IPhone 11":      690000,
	"MacBook Air":    1290000,
	"MacBook Pro 13": 1690000,
	"MacBook Pro 14": 2690000,
	"MacBook Pro 16": 3360000,
	"IMac 24":        1690000,
	"IPad Pro":       990000,
	"IPad Air":       779000,
	"IPad Mini":      649000,
	"IPad":           449000,
}

var Categories = []string{
	"IPhone", "MacBook", "IMac", "IPad",
}

// pick random promo category from list
var PromoCategory = []string{
	"Haloween", "Christmas", "New Year", "Valentine", "Easter", "Thanksgiving",
}

func GetName() string {
	return Names[rand.Intn(len(Names))]
}

func GetMailHost() string {
	return MailHost[rand.Intn(len(MailHost))]
}

func GetNumericChar(digit float64) string {
	max := int(math.Pow(10, digit))
	number := rand.Intn(max)
	return fmt.Sprintf("%04d", number)
}

func GetUUID() string {
	uuidStr := uuid.New().String()
	return uuidStr
}

func GetNickname(username string) string {
	return fmt.Sprintf("%s%s", username, GetNumericChar(4))
}

func GetEmail(nickname string) string {
	return fmt.Sprintf("%s@%s", nickname, GetMailHost())
}

func GetPassword(length int) string {
	var password []rune
	for i := 0; i < length; i++ {
		password = append(password, []rune(fmt.Sprintf("%c", rand.Intn(26)+97))[0])
	}
	return string(password)
}

func GetSHA256Password(password string) string {
	newPassword := utils.PlainTextToSHA256(password)
	return newPassword
}

func GetSex() string {
	if rand.Int()%2 == 0 {
		return "M"
	}
	return "F"
}

func GetBirth() time.Time {
	year := rand.Intn(time.Now().Year() - 2000)
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1
	return time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func GetPhone() string {
	// TODO: with +82 or other country code
	// #ref, https://countrycode.org/
	return fmt.Sprintf("+8210%08s", GetNumericChar(8))
}

func GetDiffOperator() string {
	if rand.Int()%2 == 0 {
		return ">="
	}
	return "<="
}
