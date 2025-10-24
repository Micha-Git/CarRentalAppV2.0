package operations

import (
	"fmt"
	"rentalmanagement/logic/model"
	"testing"
)

func TestRemoveRentableCarOperation(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                 string
		vin                  model.Vin
		expectedErrorMessage string
	}{
		{
			"Success",
			initialCarsInDatabase[0].Vin,
			"",
		},
		{
			"Failure",
			model.Vin{Vin: "invalidVin"},
			"Failed to remove rentable car with VIN invalidVin: Database failed to remove rentable car with VIN invalidVin: There is no rentable car matching the given condition",
		},
		{
			"Success2",
			initialCarsInDatabase[1].Vin,
			"",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := rentableCarsCollectionOperations.RemoveRentableCar(tc.vin)

			// No error expected
			if tc.expectedErrorMessage == "" && err != nil {
				t.Fatalf("Test case %v should not create error", tc.name)
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" {
				msg := fmt.Sprintf("Test case %v should produce error with message: '%v'", tc.name, tc.expectedErrorMessage)
				if err == nil {
					t.Fatalf("%s, but didn't error", msg)
				} else if err.Error() != tc.expectedErrorMessage {
					t.Fatalf("%s, but instead produces: '%v'", msg, err)
				}
			}
		})
	}
}
