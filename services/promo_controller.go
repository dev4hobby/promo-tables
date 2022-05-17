package services

import (
	"promo-tables/database/models"
	"promo-tables/utils"
	"promo-tables/utils/mock"
)

func (h *Handler) GetAllPromoControllers() (promoControllers []models.PromoController) {
	promoControllers, err := h.db.GetAllPromoControllers()
	utils.CheckError(err)
	return promoControllers
}

func (h *Handler) AddPromoControllerByPromoCategories(promoCategories []models.PromoCategory, enabled bool) error {
	promoControllers := []models.PromoController{}

	for _, promoCategory := range promoCategories {
		subscribeCondition := utils.MustMarshal(models.PromoControllerSubscribeCondition{
			Sex: mock.GetSex(),
			Birth: models.SubscribeConditionBirth{
				Timestamp: mock.GetBirth(),
				Operator:  mock.GetDiffOperator(),
			},
		})
		promoController := models.PromoController{
			PromoCategoryID:    promoCategory.ID,
			Description:        promoCategory.Name,
			BuyLimitCount:      utils.GenerateRandomNumber(1, 3),
			Enabled:            enabled,
			StartAt:            utils.GetUTCNowWithDiff(utils.GenerateRandomNumber(-24, 0)),
			EndAt:              utils.GetUTCNowWithDiff(utils.GenerateRandomNumber(18, 72)),
			SubscribeCondition: subscribeCondition,
		}
		promoControllers = append(promoControllers, promoController)
	}
	return h.db.AddPromoControllers(promoControllers)
}
