package operations

import (
	"rentalmanagement/logic/model"
	"testing"
	"time"
)

func TestListCustomerRentals(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	// make one rental for Tesla Model 3 Mannheim
	rental1, err := customerOperations.RentCar("cus0", time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 16, 12, 0, 0, 0, time.UTC), initialCarsInDatabase[0].Vin)
	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v", initialCarsInDatabase[0].Vin)
	}

	// make two rentals for Fiat 500e Mannheim
	rental2, err := customerOperations.RentCar("cus1", time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 16, 12, 0, 0, 0, time.UTC), initialCarsInDatabase[2].Vin)
	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v", initialCarsInDatabase[2].Vin)
	}

	rental3, err := customerOperations.RentCar("cus1", time.Date(2024, time.January, 25, 12, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 26, 12, 0, 0, 0, time.UTC), initialCarsInDatabase[2].Vin)
	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v", initialCarsInDatabase[2].Vin)
	}

	cases := []struct {
		name                 string
		customerId           string
		expectedErrorMessage string
		expectedRentals      []model.Rental
	}{
		{
			name:                 "Empty CustomerID",
			customerId:           "",
			expectedErrorMessage: "Customer ID can't be empty",
			expectedRentals:      nil,
		},
		{
			name:                 "One Rental",
			customerId:           "cus0",
			expectedErrorMessage: "",
			expectedRentals:      []model.Rental{rental1},
		},
		{
			name:                 "No Rental",
			customerId:           "cus2",
			expectedErrorMessage: "",
			expectedRentals:      nil,
		},
		{
			name:                 "Two Rentals",
			customerId:           "cus1",
			expectedErrorMessage: "",
			expectedRentals:      []model.Rental{rental2, rental3},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := rentalsCollectionOperations.ListCustomerRentals(tc.customerId)

			// No error expected
			if tc.expectedErrorMessage == "" {
				if err != nil {
					t.Fatalf("Test case %v should not create error", tc.name)
				} else if !rentalArraysEqual(tc.expectedRentals, result) {
					t.Fatalf("Test case '%v' should have a different result", tc.name)
				}
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" && (err == nil || err.Error() != tc.expectedErrorMessage) {
				t.Fatalf("Test case %v should produce error with message: %v, but instead produces: %v", tc.name, tc.expectedErrorMessage, err.Error())
			}
		})
	}
}
