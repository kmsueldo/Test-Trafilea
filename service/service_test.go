package service

import (
	"challenge-trafilea/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateAndSaveOrder(t *testing.T) {
	var tests = []struct {
		name                string
		cartDataMock        models.Cart
		orderParam          models.Order
		expectedOrderResult models.Order
		expectedError       bool
	}{
		{
			name: "CalculateAndSaveOrder return Order created ok when an extra coffee is given",
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
				},
			},
			orderParam: models.Order{
				CartID: "Cart1",
			},
			expectedOrderResult: models.Order{
				CartID: "Cart1",
				Totals: models.Totals{
					Products: []models.Product{
						{
							ProductID: "1",
							Name:      "coffee",
							Category:  "Coffee",
							Price:     4,
							Quantity:  2,
						},
						{
							ProductID: "CoffeeFree1",
							Name:      "coffee",
							Category:  "Coffee",
							Price:     0,
							Quantity:  1,
						},
					},
					AmountShipping:  20,
					PercentDiscount: 0,
					AmountOrder:     4,
				},
			},
			expectedError: false,
		},
		{
			name: "CalculateAndSaveOrder return Order created ok when is shipping free",
			cartDataMock: models.Cart{
				CartID: "Cart1",
				UserID: "User1",
				Products: []models.Product{
					{
						ProductID: "2",
						Name:      "Table",
						Category:  "Equipment",
						Price:     40,
						Quantity:  4,
					},
				},
			},
			orderParam: models.Order{
				CartID: "Cart1",
			},
			expectedOrderResult: models.Order{
				CartID: "Cart1",
				Totals: models.Totals{
					Products: []models.Product{
						{
							ProductID: "2",
							Name:      "Table",
							Category:  "Equipment",
							Price:     40,
							Quantity:  4,
						},
					},
					AmountShipping:  0,
					PercentDiscount: 0,
					AmountOrder:     40,
				},
			},
			expectedError: false,
		},
		{
			name: "CalculateAndSaveOrder return Order created ok when get 10% discount",
			cartDataMock: models.Cart{
				CartID: "Cart1",
				UserID: "User1",
				Products: []models.Product{
					{
						ProductID: "3",
						Name:      "Pencil",
						Category:  "Accessories",
						Price:     80,
						Quantity:  10,
					},
				},
			},
			orderParam: models.Order{
				CartID: "Cart1",
			},
			expectedOrderResult: models.Order{
				CartID: "Cart1",
				Totals: models.Totals{
					Products: []models.Product{
						{
							ProductID: "3",
							Name:      "Pencil",
							Category:  "Accessories",
							Price:     80,
							Quantity:  10,
						},
					},
					AmountShipping:  20,
					PercentDiscount: 10,
					AmountOrder:     80,
				},
			},
			expectedError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := SaveCart(tt.cartDataMock)
			if err != nil {
				t.Errorf("failed to add cart mock %v", err.Error())
			}
			orderCalculated, err := CalculateAndSaveOrder(tt.orderParam)
			assert.Equal(t, orderCalculated, tt.expectedOrderResult)
			assert.Equal(t, tt.expectedError, err != nil)

		})
	}

}
