package test

import (
	"promo-tables/services"
	"promo-tables/utils"
)

func GetServiceHandler() services.HandlerInterface {
	handler, err := services.NewHandler()
	utils.CheckError(err)
	return handler
}
