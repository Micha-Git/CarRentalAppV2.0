package operations

import (
	"fleetmanagement/logic/model"
	"testing"
)

func TestViewCarInformation(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                   string
		vin                    model.Vin
		expectedCarInformation model.Car
		expectedErrorMessage   string
	}{
		/*{
			name:                   "Vin not Valid",
			vin:                    model.Vin{Vin: "invalidVin"},
			expectedCarInformation: model.Car{},
			expectedErrorMessage:   "Failed to retrieve car with vin",
		},*/
		{
			name:                   "Success",
			vin:                    carRepoCars[0].Vin,
			expectedCarInformation: carRepoCars[0],
			expectedErrorMessage:   "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := carOperations.ViewCarInformation(tc.vin)

			// Error unexpected
			if tc.expectedErrorMessage == "" {
				if err != nil {
					t.Fatalf("Test case %v should not create error", tc.name)
				} else if result.Vin != tc.vin {
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
