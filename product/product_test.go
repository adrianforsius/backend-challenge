// We test using a different package to make sure they don't rely on internal state
package product_test

import (
	"testing"

	"github.com/adrianforsius/backend-challenge/product"
)

func TestDiscountSmallPurchase(t *testing.T) {
	merchendise := product.Merchendise
	var products []product.Product
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[0]) // pen
	}
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[1]) // tshirt
	}
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[2]) // mug
	}

	total := product.Discont(products)
	expected := 3250
	if total != expected {
		t.Errorf("got = %d; want %d", total, expected)
	}
}

// I don't udnerstand the spec and I can't guess what this is suppose to be
func TestDiscountSmallExtraPenPurchase(t *testing.T) {
	merchendise := product.Merchendise
	var products []product.Product
	for i := 0; i < 2; i++ {
		products = append(products, merchendise[0]) // pen
	}
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[1]) // tshirt
	}

	total := product.Discont(products)
	expected := 2500
	if total != expected {
		t.Errorf("got = %d; want %d", total, expected)
	}
}

// This test is working by contrary to what the spec say, IE I switched pen and t-shirt discount
func TestDiscountTShirtPurchase(t *testing.T) {
	merchendise := product.Merchendise
	var products []product.Product
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[0]) // pen
	}
	for i := 0; i < 4; i++ {
		products = append(products, merchendise[1]) // tshirt
	}

	total := product.Discont(products)
	expected := 6500
	if total != expected {
		t.Errorf("got = %d; want %d", total, expected)
	}
}

func TestDiscountExtraPensAndTShirtPurchase(t *testing.T) {
	merchendise := product.Merchendise
	var products []product.Product
	for i := 0; i < 3; i++ {
		products = append(products, merchendise[0]) // pen
	}
	for i := 0; i < 3; i++ {
		products = append(products, merchendise[1]) // tshirt
	}
	for i := 0; i < 1; i++ {
		products = append(products, merchendise[2]) // mug
	}

	total := product.Discont(products)
	expected := 6250
	if total != expected {
		t.Errorf("got = %d; want %d", total, expected)
	}
}
