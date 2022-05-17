package main

import (
	"log"
	"math/rand"
	"promo-tables/services"
	"promo-tables/utils"
	"time"
)

func main() {
	log.Print("[SYS] Hello")
	rand.Seed(time.Now().UnixNano())
	cmd := utils.GetFlag()
	services.MainHandler(cmd)
	log.Print("[SYS] Bye")
}
