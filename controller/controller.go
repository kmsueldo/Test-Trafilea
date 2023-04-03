package controller

import (
	"challenge-trafilea/service"
	"encoding/json"
	"log"
	"net/http"
)

const (
	_emptyString = ""
)

// CreateCart create a cart with the information obtained from the payload
func CreateCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	newCart, err := service.GetCartPayloadFromRequest(request)
	if err != nil {
		log.Println("Failed getting payload from create cart")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Failed getting payload from create cart")
		return
	}

	if err = service.SaveCart(newCart); err != nil {
		log.Println("Failed saving cart in storage")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Failed saving cart in storage")
		return
	}
	log.Println("The cart has been created", newCart)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(newCart)
	return
}

// GetCart get specific cart information
func GetCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	cart := service.GetCart()
	if cart.UserID == _emptyString {
		log.Println("The cart are empty")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode("The cart are empty")
		return
	}

	log.Println("The cart obtained", cart)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cart)
	return
}

// AddProductToCart add a product to cart
func AddProductToCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	newProduct, err := service.GetProductPayloadFromRequest(request)
	if err != nil {
		log.Println("Failed getting payload from add product")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Failed getting payload from add product")
		return
	}

	if err = service.UpdateCartWithNewProduct(newProduct); err != nil {
		log.Println("Failed saving product in cart")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Failed saving product in cart")
		return
	}

	log.Println("The product has been added", newProduct)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(newProduct)
	return
}

// UpdateQuantityProductInTheCart update the quantity of a specific product in the cart
func UpdateQuantityProductInTheCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	productToUpdate, err := service.GetProductPayloadFromRequest(request)
	if err != nil {
		log.Println("Failed getting payload from update product")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Failed getting payload from update product")
		return
	}

	if err = service.UpdateQuantityProductToCart(productToUpdate); err != nil {
		log.Println("Failed updating quantity product in cart")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Failed updating quantity product in cart")
		return
	}

	log.Println("The quantity of the product has been updated", productToUpdate)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(productToUpdate)
	return
}

// CreateOrder create a cart order
func CreateOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	orderID, err := service.GetOrderPayloadFromRequest(request)
	if err != nil {
		log.Println("Failed getting payload from create order")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Failed getting payload from create order")
		return
	}

	orderSaved, err := service.CalculateAndSaveOrder(orderID)
	if err != nil {
		log.Println("Failed save and calculate order")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Failed create order")
		return
	}

	log.Println("The order has been saved", orderSaved)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(orderSaved)
	return
}
