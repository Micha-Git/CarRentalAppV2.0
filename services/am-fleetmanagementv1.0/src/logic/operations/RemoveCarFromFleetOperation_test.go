package operations

import (
	"fleetmanagement/logic/model"
	"testing"
)

func TestRemoveCarFromFleet(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                 string
		vin                  model.Vin
		expectedResult       bool
		expectedErrorMessage string
	}{
		{
			name:                 "Vin not Valid",
			vin:                  model.Vin{Vin: "invalidVin"},
			expectedResult:       false,
			expectedErrorMessage: "Car with ID {invalidVin} does not exist: Database failed to find car with vin invalidVin: record not found",
		},
		{
			name:                 "Success",
			vin:                  carRepoCars[0].Vin,
			expectedResult:       true,
			expectedErrorMessage: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := fleetOperations.RemoveCarFromFleet(tc.vin)

			// Check if car was removed
			if tc.expectedErrorMessage == "" && err == nil {
				_, err2 := postgresRepo.GetCar(tc.vin)
				if err2 == nil {
					t.Fatalf("Test case \"%v\" executed as expected without an error message, but didn't remove car from fleet repository", tc.name)
				}
			}

			// Error unexpected
			if tc.expectedErrorMessage == "" {
				if err != nil {
					t.Fatalf("Test case %v should not create error: %v", tc.name, err.Error())
				} else if tc.expectedResult != result {
					t.Fatalf("Test case '%v' should return different result", tc.name)
				}
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" && (err == nil || err.Error() != tc.expectedErrorMessage) {
				t.Fatalf("Test case %v should produce error with message: %v, but instead produces: %v", tc.name, tc.expectedErrorMessage, err.Error())
			}
		})
	}

}
