package repository

import (
	"challenge-trafilea/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateQuantityProductToCart(t *testing.T) {
	var tests = []struct {
		name            string
		cartDataMock    models.Cart
		productToUpdate models.Product
		expectedError   bool
	}{
		{
			name: "UpdateQuantityProductToCart is ok when find the product to update",
			cartDataMock: models.Cart{
				CartID: "Cart1",
				UserID: "User1",
				Products: []models.Product{
					{
						ProductID: "1",
						Name:      "coffee",
						Category:  "Coffee",
						Price:     4,
						Quantity:  2,
					},
					{
						ProductID: "2",
						Name:      "table",
						Category:  "Equipment",
						Price:     20,
						Quantity:  2,
					},
				},
			},
			productToUpdate: models.Product{
				ProductID: "1",
				Name:      "coffee",
				Category:  "Coffee",
				Price:     6,
				Quantity:  3,
			},
			expectedError: false,
		},
		{
			name: "UpdateQuantityProductToCart is ok when can not find the product to update",
			cartDataMock: models.Cart{
				CartID: "Cart1",
				UserID: "User1",
				Products: []models.Product{
					{
						ProductID: "1",
						Name:      "coffee",
						Category:  "Coffee",
						Price:     4,
						Quantity:  2,
					},
					{
						ProductID: "2",
						Name:      "table",
						Category:  "Equipment",
						Price:     20,
						Quantity:  2,
					},
				},
			},
			productToUpdate: models.Product{
				ProductID: "3",
				Name:      "coffee",
				Category:  "Coffee",
				Price:     6,
				Quantity:  3,
			},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := SaveCartInStore(tt.cartDataMock)
			if err != nil {
				t.Errorf("failed to add cart mock %v", err.Error())
			}
			orderCalculated := UpdateQuantityProductToCart(tt.productToUpdate)
			assert.Equal(t, tt.expectedError, orderCalculated != nil)

		})
	}

}
