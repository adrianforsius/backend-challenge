package product

import (
	"fmt"
	"math"
	"sync"

	"github.com/google/uuid"
)

type Product struct {
	Price int    `json:"price"` // Reprecents cents
	Name  string `json:"name"`
	Code  string `json:"code"`
}

type BasketStorage struct {
	baskets map[string][]Product
	l       *sync.Mutex
}

func NewBasket() BasketStorage {
	return BasketStorage{
		baskets: make(map[string][]Product, 0),
		l:       &sync.Mutex{},
	}
}

func (b *BasketStorage) Get(id string) ([]Product, error) {
	b.l.Lock()
	defer b.l.Unlock()
	return Basket(b.baskets, id)
}

func (b *BasketStorage) New() string {
	b.l.Lock()
	defer b.l.Unlock()
	id := uuid.New().String()
	b.baskets[id] = make([]Product, 0)
	return id
}

func (b *BasketStorage) Add(products []Product, id string) ([]Product, error) {
	b.l.Lock()
	defer b.l.Unlock()
	if _, ok := b.baskets[id]; !ok {
		return nil, fmt.Errorf("no such basket found")
	}
	b.baskets[id] = append(b.baskets[id], products...)
	return b.baskets[id], nil
}

func (b *BasketStorage) Remove(id string) error {
	b.l.Lock()
	defer b.l.Unlock()
	if _, ok := b.baskets[id]; !ok {
		return fmt.Errorf("no such basket found")
	}
	delete(b.baskets, id)
	return nil
}

func Basket(baskets map[string][]Product, ID string) ([]Product, error) {
	for id, basket := range baskets {
		if id == ID {
			return basket, nil
		}
	}
	return nil, fmt.Errorf("No such basket found in baskets")
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

	if items["TSHIRT"] >= 3 {
		penDiscounts := int(math.Trunc(float64(items["TSHIRT"] / 3)))
		total -= penDiscounts * PRICE_TSHIRT
	}

	if items["PEN"] >= 3 {
		discount := int(math.Trunc(float64(items["PEN"]) * 25 / 100))
		total -= discount
	}
	return total
}
