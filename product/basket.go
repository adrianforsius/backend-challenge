package product

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

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
