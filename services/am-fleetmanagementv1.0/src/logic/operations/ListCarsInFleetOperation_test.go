package operations

import (
	"fleetmanagement/logic/model"
	"testing"
)

func TestListCarsInFleet(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                 string
		fleetId              string
		expectedCars         []model.Car
		expectedErrorMessage string
	}{
		{
			name:                 "FleetId Not Valid",
			fleetId:              "InvalidFleetId",
			expectedCars:         nil,
			expectedErrorMessage: "Fleet with ID InvalidFleetId does not exist: Database failed to get fleet with ID InvalidFleetId: record not found",
		},
		{
			name:                 "No Car in fleet",
			fleetId:              "1",
			expectedCars:         nil,
			expectedErrorMessage: "",
		},
		{
			name:                 "One Car in fleet",
			fleetId:              "1",
			expectedCars:         []model.Car{carRepoCars[0]},
			expectedErrorMessage: "",
		},
		{
			name:                 "Two Cars in fleet",
			fleetId:              "1",
			expectedCars:         []model.Car{carRepoCars[0], carRepoCars[1]},
			expectedErrorMessage: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := fleetOperations.ListCarsInFleet(tc.fleetId)

			// Error unexpected
			if tc.expectedErrorMessage == "" {
				if err != nil {
					t.Fatalf("Test case %v should not create error", tc.name)
				}
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" && (err == nil || err.Error() != tc.expectedErrorMessage) {
				t.Fatalf("Test case %v should produce error with message: %v, but instead produces: %v", tc.name, tc.expectedErrorMessage, err.Error())
			}
		})
	}

}
