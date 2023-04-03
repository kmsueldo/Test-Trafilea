package service

import (
	"challenge-trafilea/models"
	"challenge-trafilea/repository"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	_emptyString = ""
	_coffee      = "Coffee"
	_equipment   = "Equipment"
	_accessories = "Accessories"
)

var ErrCartIdEmptyOrNotExist = errors.New("cartID is empty or not exist")

func GetCartPayloadFromRequest(request *http.Request) (models.Cart, error) {
	var newCart models.Cart

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return models.Cart{}, err
	}

	err = json.Unmarshal(reqBody, &newCart)
	if err != nil {
		return models.Cart{}, err
	}

	return newCart, nil
}

func SaveCart(cart models.Cart) error {
	if err := repository.SaveCartInStore(cart); err != nil {
		return err
	}
	return nil
}

func GetCart() models.Cart {
	return repository.SelectCart()

}

func GetProductPayloadFromRequest(request *http.Request) (models.Product, error) {
	var product models.Product

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return models.Product{}, err
	}

	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func GetOrderPayloadFromRequest(request *http.Request) (models.Order, error) {
	var order models.Order

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return models.Order{}, err
	}

	err = json.Unmarshal(reqBody, &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func UpdateCartWithNewProduct(newProduct models.Product) error {
	if err := repository.UpdateCart(newProduct); err != nil {
		return err
	}
	return nil
}

func UpdateQuantityProductToCart(productToUpdate models.Product) error {
	if err := repository.UpdateQuantityProductToCart(productToUpdate); err != nil {
		return err
	}
	return nil
}

// CalculateAndSaveOrder calculate the special rules for coffee, shipping and discounts and save the order
func CalculateAndSaveOrder(order models.Order) (models.Order, error) {

	if order.CartID != _emptyString && order.CartID == repository.Cart.CartID {
		order.Totals.Products = evaluateCategoryForFreeCoffee(order, repository.Cart.Products)
		order.Totals.AmountShipping = evaluateCategoryForFreeShipping(order.Totals.Products)
		order.Totals.PercentDiscount = evaluateCategoryForDiscounts(order.Totals.Products)
		order.Totals.AmountOrder = calculateAmountTotalOrder(order.Totals.Products)

		if err := repository.SaveOrderInStore(order); err != nil {
			return models.Order{}, err
		}
	} else {
		log.Println(ErrCartIdEmptyOrNotExist)
		return order, ErrCartIdEmptyOrNotExist
	}
	return order, nil
}

// evaluateCategoryForFreeCoffee evaluate coffee category products and add an extra free coffee
// to the order when there are more than 2 coffee categories
func evaluateCategoryForFreeCoffee(order models.Order, products []models.Product) []models.Product {
	var countCategoryCoffee int

	for index := range products {
		if products[index].Category == _coffee {
			countCategoryCoffee = countCategoryCoffee + products[index].Quantity
		}
		if countCategoryCoffee >= 2 {
			var newCoffeeFree = models.Product{
				ProductID: "CoffeeFree1",
				Name:      "coffee",
				Category:  "Coffee",
				Price:     0,
				Quantity:  1,
			}
			products = append(products, newCoffeeFree)
			order.Totals.Products = products
			return products
		}
	}
	return products
}

// evaluateCategoryForFreeShipping evaluate the products of the equipment category and when
// there are more than 3 products the shipping is free
func evaluateCategoryForFreeShipping(products []models.Product) int {
	var countCategoryEquipment int

	costShippingFree := 0
	generalShippingCost := 20
	for index := range products {
		if products[index].Category == _equipment {
			countCategoryEquipment++
			if products[index].Quantity > 3 {
				return costShippingFree
			}
		}
		if countCategoryEquipment > 3 {
			return costShippingFree
		}
	}
	return generalShippingCost
}

// evaluateCategoryForDiscounts evaluate the accessories category products and when the expense
// exceeds 70 dollars a 10% discount is applied
func evaluateCategoryForDiscounts(products []models.Product) int {
	var amountDiscount int
	var expenseCategoryAccessories float32

	for index := range products {
		if products[index].Category == _accessories {
			expenseCategoryAccessories = expenseCategoryAccessories + products[index].Price
		}
		if expenseCategoryAccessories > 70 {
			amountDiscount = 10
			return amountDiscount
		}
	}
	return amountDiscount
}

// calculateAmountTotalOrder calculates the total amount of the order
func calculateAmountTotalOrder(products []models.Product) float32 {
	var amountTotalOrder float32

	for index := range products {
		amountTotalOrder = amountTotalOrder + products[index].Price
	}
	return amountTotalOrder

}
