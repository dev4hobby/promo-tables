package services

import (
	"encoding/json"
	"promo-tables/database/models"
	"promo-tables/utils"
)

func (h *Handler) AddUserStoreSubscriptions() error {

	promoControllers, err := h.db.GetAllPromoControllers()
	utils.CheckError(err)
	userSubscriptionMap := map[int][]int{}
	for _, o := range promoControllers {
		var subscriptionCondition models.PromoControllerSubscribeCondition
		json.Unmarshal([]byte(o.SubscribeCondition), &subscriptionCondition)

		users := h.db.GetUsersFilteredBySubscribeCondition(subscriptionCondition)

		for _, user := range users {
			userSubscriptionMap[user.ID] = append(userSubscriptionMap[user.ID], o.ID)
		}
	}

	subscription := []models.UserStoreSubscription{}
	for userID, promoControllerIDList := range userSubscriptionMap {
		subscription = append(subscription, models.UserStoreSubscription{
			UserID:                userID,
			PromoControllerIDList: utils.MustMarshal(promoControllerIDList),
		})
	}

	return h.db.AddUserStoreSubscriptions(subscription)
}
