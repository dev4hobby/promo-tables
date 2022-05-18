package test

import (
	"promo-tables/services"
	"promo-tables/utils"
	"testing"
)

func GetServiceHandler(t *testing.T) services.HandlerInterface {
	// if err != nil; retry
	handler, err := services.NewHandler()
	retryMaxCount := 6
	tried := 0
	for err != nil {
		handler, err = services.NewHandler()
		utils.SleepSeconds(5)
		tried++
		if tried >= retryMaxCount {
			t.Errorf("[FAIL] Failed to Initialize Database")
			panic(err)
		}
		t.Log("[DB] Re-connecting... ")
	}
	return handler
}
