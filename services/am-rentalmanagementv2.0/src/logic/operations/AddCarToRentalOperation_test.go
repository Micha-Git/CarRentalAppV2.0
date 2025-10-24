package operations

import (
	"fmt"
	"rentalmanagement/logic/model"
	"testing"
)

func TestAddCarToRentalOperation(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                 string
		vin                  model.Vin
		location             string
		expectedErrorMessage string
	}{
		{
			name:                 "Success",
			vin:                  carRepoCars[0].Vin,
			location:             "Karlsruhe",
			expectedErrorMessage: "",
		},
		{
			name:                 "Vin Not Valid",
			vin:                  model.Vin{Vin: "invalidVin"},
			location:             "Mannheim",
			expectedErrorMessage: "Failed to get information about car with VIN invalidVin: API request failed with status code: 400",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := rentableCarsCollectionOperations.AddCarToRental(tc.vin, tc.location)

			// Check if car was added
			if tc.expectedErrorMessage == "" && err == nil {
				_, err = rentableCarRepository.GetRentableCar(tc.vin)
				if err != nil {
					t.Fatalf("Test case \"%v\" executed as expected without an error message, but didn't add car to rental repository", tc.name)
				}
			}

			// No error expected
			if tc.expectedErrorMessage == "" && err != nil {
				t.Fatalf("Test case \"%v\" should not create error, but instead creates error: \"%v\"", tc.name, err.Error())
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" {
				msg := fmt.Sprintf("Test case \"%v\" should produce error with message: \"%v\"", tc.name, tc.expectedErrorMessage)
				if err == nil {
					t.Fatalf("%s, but didn't error", msg)
				} else if err.Error() != tc.expectedErrorMessage {
					t.Fatalf("%s, but instead produces: '%v'", msg, err)
				}
			}
		})
	}
}
