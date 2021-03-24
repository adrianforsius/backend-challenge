package product

import (
	"fmt"
	"math"
)

type Product struct {
	Price int    `json:"price"` // Reprecents cents
	Name  string `json:"name"`
	Code  string `json:"code"`
}

const (
	PRICE_PEN    = 500
	PRICE_TSHIRT = 2000
	PRICE_MUG    = 750
)

var Merchendise = []Product{
	{
		Code:  "PEN",
		Name:  "Lana pen",
		Price: PRICE_PEN,
	},
	{
		Code:  "TSHIRT",
		Name:  "Lana T-Shirt",
		Price: PRICE_TSHIRT,
	},
	{
		Code:  "MUG",
		Name:  "Lana Coffee Mug",
		Price: PRICE_MUG,
	},
}

func Validate(code string) (Product, error) {
	for _, m := range Merchendise {
		if code == m.Code {
			return m, nil
		}
	}
	return Product{}, fmt.Errorf("no such code")
}

func Discont(products []Product) int {
	items := map[string]int{
		"PEN":    0,
		"TSHIRT": 0,
		"MUG":    0,
	}
	var total int
	for _, p := range products {
		items[p.Code] += 1
		total += p.Price
	}

	if items["PEN"] >= 2 {
		penDiscounts := int(math.Trunc(float64(items["PEN"] / 2)))
		total -= penDiscounts * PRICE_PEN
	}

	if items["TSHIRT"] >= 3 {
		discount := int(math.Round(float64(items["TSHIRT"]*PRICE_TSHIRT) * 25 / 100))
		total -= discount
	}
	return total
}
