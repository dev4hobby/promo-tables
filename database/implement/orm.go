package implement

import (
	"promo-tables/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ORM struct {
	*gorm.DB
}

func NewORM() (*ORM, error) {
	env := utils.GetEnv()
	dsn := env.MySQLUser + ":" + env.MySQLPassword + "@tcp(" + env.MySQLHost + ":" + env.MySQLPort + ")/" + env.MySQLDatabase + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &ORM{
		DB: db,
	}, err
}
