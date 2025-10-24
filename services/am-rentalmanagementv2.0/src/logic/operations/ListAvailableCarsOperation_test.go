package operations

import (
	"rentalmanagement/logic/model"
	"testing"
	"time"
)

func TestListAvailableCars(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cases := []struct {
		name                 string
		startDate            time.Time
		endDate              time.Time
		location             string
		expectedCars         []model.Vin
		expectedErrorMessage string
	}{
		{
			name:                 "Invalid Dates: startDate not before endDate",
			startDate:            time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			endDate:              time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			location:             "Karlsruhe",
			expectedCars:         []model.Vin{},
			expectedErrorMessage: "StartDate must be before EndDate",
		},
		{
			name:      "One available rentable car",
			startDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
			location:  "Mannheim",
			expectedCars: []model.Vin{
				initialCarsInDatabase[4].Vin, // VW ID.2 Mannheim
			},
			expectedErrorMessage: "",
		},
		{
			name:      "Tesla Model 3 Karlsruhe available again",
			startDate: time.Date(2024, time.December, 1, 0, 0, 0, 1, time.UTC),
			endDate:   time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC),
			location:  "Karlsruhe",
			expectedCars: []model.Vin{
				initialCarsInDatabase[0].Vin, // Tesla Model 3 Karlsruhe
				initialCarsInDatabase[1].Vin, // Seat Leon Karlsruhe
				initialCarsInDatabase[2].Vin, // Fiat 500e Karlsruhe
				initialCarsInDatabase[3].Vin, // Audi A3 Karlsruhe
			},
			expectedErrorMessage: "",
		},
		{
			name:      "Three available rentable cars",
			startDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
			location:  "Karlsruhe",
			expectedCars: []model.Vin{
				initialCarsInDatabase[1].Vin, // Seat Leon Karlsruhe
				initialCarsInDatabase[2].Vin, // Fiat 500e Karlsruhe
				initialCarsInDatabase[3].Vin, // Audi A3 Karlsruhe
			},
			expectedErrorMessage: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t, customerOperations)
			defer teardownSubTest(t)

			result, err := rentalsCollectionOperations.ListAvailableCars(tc.startDate, tc.endDate, tc.location)

			// Expected error mismatch
			if tc.expectedErrorMessage != "" &&
				(result != nil || err == nil || tc.expectedErrorMessage != err.Error()) {
				t.Fatalf("Test case '%v' should produce error with message: '%v'", tc.name, tc.expectedErrorMessage)
			}

			// Unexpected error
			if tc.expectedErrorMessage == "" {
				if err != nil {
					t.Fatalf("Test case '%v' should not create error", tc.name)
				} else if result == nil || !carsHaveExpectedVins(tc.expectedCars, result) {
					t.Fatalf("Test case '%v' should have a different result", tc.name)
				}
			}
		})
	}
}

func setupSubTest(t *testing.T, customerOperations CustomerOperations) func(t *testing.T) {
	t.Logf("Setting up sub test %s", t.Name())

	rental, err := customerOperations.RentCar(
		"1",
		time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
		initialCarsInDatabase[0].Vin, // Tesla Model 3 Karlsruhe
	)

	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v", initialCarsInDatabase[0].Vin)
	}

	return func(t *testing.T) {
		t.Log("teardown sub test")

		err := customerOperations.CancelRental(rental.CustomerId, rental.Id)

		if err != nil {
			t.Fatalf("Failed to cancel rental with id '%v", rental.Id)
		}
	}
}
