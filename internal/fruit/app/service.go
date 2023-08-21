package app

import "gitlab.cmpayments.local/creditcard/fruit-price-calculator/internal/entities"

type fruitService struct {
}

func NewFruitService() *fruitService {
	return &fruitService{}
}

func (s fruitService) CalculateAveragePrice(fruits entities.Fruits) entities.Fruits {
	var (
		fruitPrices   = make(map[string]result, 0)
		fruitAverages = entities.Fruits{}
	)

	for _, fruit := range fruits {
		r := fruitPrices[fruit.Type]
		fruitPrices[fruit.Type] = result{total: r.total + fruit.Price, quantity: r.quantity + 1}
	}

	for key, result := range fruitPrices {
		fruitAverages = append(fruitAverages, entities.Fruit{Type: key, Price: result.total / float64(result.quantity)})
	}

	return fruitAverages
}

type result struct {
	total    float64
	quantity int
}
