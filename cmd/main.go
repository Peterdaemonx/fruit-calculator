package main

import (
	fruitApp "gitlab.cmpayments.local/creditcard/fruit-price-calculator/internal/fruit/app"
	"gitlab.cmpayments.local/creditcard/fruit-price-calculator/internal/fruit/web"
	"log"
	"os"
)

func main() {

	handler := web.NewFruitHandler(fruitApp.NewFruitService())
	handler.PrintFruitsAverageFromFile(os.Args)

	defer func() {
		if err := recover(); err != nil {
			log.Println("unexpected application error:", err)
		}
	}()
}
