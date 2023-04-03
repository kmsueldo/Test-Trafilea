package repository

import (
	"challenge-trafilea/models"
	"errors"
	"log"
)

var (
	Cart                     models.Cart
	Order                    models.Order
	ErrFailedToAddCart       = errors.New("failed to add cart")
	ErrFailedToAddProduct    = errors.New("failed to add product")
	ErrFailedProductNotExist = errors.New("product not found")
	ErrFailedToCreateOrder   = errors.New("failed to create order")
)

const (
	_emptyString = ""
)

func SaveCartInStore(cart models.Cart) error {

	Cart = cart
	if Cart.UserID == _emptyString {
		log.Println(ErrFailedToAddCart)
		return ErrFailedToAddCart
	}

	return nil
}

func SelectCart() models.Cart {
	return Cart
}

func UpdateCart(newProduct models.Product) error {

	lenOriginalProductCart := len(Cart.Products)

	Cart.Products = append(Cart.Products, newProduct)

	if len(Cart.Products) <= lenOriginalProductCart {
		log.Println(ErrFailedToAddProduct)
		return ErrFailedToAddProduct
	}

	return nil
}

func UpdateQuantityProductToCart(productToUpdate models.Product) error {

	isProductUpdated := false
	for index := range Cart.Products {
		if Cart.Products[index].ProductID == productToUpdate.ProductID {
			Cart.Products[index] = productToUpdate
			isProductUpdated = true
		}
	}

	if !isProductUpdated {
		log.Println(ErrFailedProductNotExist)
		return ErrFailedProductNotExist
	}

	return nil
}

func SaveOrderInStore(order models.Order) error {

	Order = order
	if Order.CartID == _emptyString {
		log.Println(ErrFailedToCreateOrder)
		return ErrFailedToCreateOrder
	}
	return nil
}
