package product

import (
	"fmt"
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

func (b *BasketStorage) Add(products []Product) string {
	b.l.Lock()
	defer b.l.Unlock()
	id := uuid.New().String()
	b.baskets[id] = products
	return id
}

func Basket(baskets map[string][]Product, ID string) ([]Product, error) {
	for id, basket := range baskets {
		if id == ID {
			return basket, nil
		}
	}
	return nil, fmt.Errorf("No such basket found in baskets")
}

var codes = []string{
	"PEN",
	"TSHIRT",
	"MUG",
}

func Validate(product Product) error {
	for _, code := range codes {
		if product.Code == code {
			return nil
		}
	}
	return fmt.Errorf("no such code")
}
