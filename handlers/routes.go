package handlers

import (
	"challenge-trafilea/controller"

	"github.com/gorilla/mux"
)

func RoutesMapper() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Methods GET
	router.HandleFunc("/api/v1/cart", controller.GetCart).Methods("GET")

	//Methods POST
	router.HandleFunc("/api/v1/create_cart", controller.CreateCart).Methods("POST")
	router.HandleFunc("/api/v1/add_product_cart", controller.AddProductToCart).Methods("POST")
	router.HandleFunc("/api/v1/create_order", controller.CreateOrder).Methods("POST")

	//Methods UPDATE
	router.HandleFunc("/api/v1/update_quantity_product", controller.UpdateQuantityProductInTheCart).Methods("PUT")

	return router

}
