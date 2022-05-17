package services

import (
	"log"
	"promo-tables/database/implement"
	"promo-tables/database/models"
	"promo-tables/utils"
)

type Handler struct {
	db implement.DBLayer
}

type HandlerInterface interface {
	GetProductByID(int) models.Product
	GetCategories() []models.Category
	GetAllProducts() []models.Product
	GetCategoriesByName(string) []models.Category
	GetAllPromoControllers() []models.PromoController
	AddUsers(int) error
	AddProducts() error
	AddCategories() error
	AddPromoControllerByPromoCategories([]models.PromoCategory, bool) error
	AddUserStoreSubscriptions() error

	InitMockTable()
	FlushAllProgress()
}

func MainHandler(cmd utils.Commands) {
	h, err := NewHandler()
	utils.CheckError(err)

	switch {
	case *cmd.InitMockTable:
		log.Print("[SYS] InitMockTable")
		h.InitMockTable()
	case *cmd.FlushAllProgress:
		log.Print("[SYS] FlushAllProgress")
		h.FlushAllProgress()
	case *cmd.HTTPServer:
		log.Print("[SYS] HTTPServer preparing")
		// if U wanna try HTTP Server
		// please implement and put it here
	}
}

func NewHandler() (HandlerInterface, error) {
	db, err := implement.NewORM()
	return &Handler{db: db}, err
}

func Hello() string {
	return "Hello, World"
}
