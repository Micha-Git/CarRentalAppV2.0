package operations

import (
	"rentalmanagement/logic/model"
	"testing"
	"time"
)

func TestRentCar(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	invalidVin := model.Vin{Vin: "invalidVin"}

	_, err := customerOperations.RentCar("0", time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 16, 12, 0, 0, 0, time.UTC), initialCarsInDatabase[0].Vin) // Tesla Model 3 Mannheim
	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v", initialCarsInDatabase[0].Vin)
	}

	cases := []struct {
		name                 string
		start                time.Time
		end                  time.Time
		vin                  model.Vin
		expectedPrice        float32
		expectedErrorMessage string
	}{
		{
			"StartAndEndDateTheSame",
			time.Date(2024, time.January, 6, 17, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 17, 0, 0, 0, time.UTC),
			initialCarsInDatabase[0].Vin,
			0,
			"StartDate must be before EndDate",
		},
		{
			"StartDateAfterEndDate",
			time.Date(2024, time.January, 7, 17, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 17, 0, 0, 0, time.UTC),
			initialCarsInDatabase[0].Vin,
			0,
			"StartDate must be before EndDate",
		},
		{
			"CarDoesNotExist",
			time.Date(2024, time.January, 5, 17, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 17, 0, 0, 0, time.UTC),
			invalidVin,
			0,
			"Car with VIN {invalidVin} does not exist or error occured checking car existence: Database failed to find rentable car with VIN {invalidVin}: record not found",
		},
		{
			"SuccessOneDay",
			time.Date(2024, time.January, 5, 17, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 17, 0, 0, 0, time.UTC),
			initialCarsInDatabase[0].Vin,
			50,
			"",
		},
		{
			"SuccessFourDays",
			time.Date(2024, time.January, 6, 18, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 10, 18, 0, 0, 0, time.UTC),
			initialCarsInDatabase[0].Vin,
			200,
			"",
		},
		{
			"CarNotAvailable",
			time.Date(2024, time.January, 15, 17, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 20, 17, 0, 0, 0, time.UTC),
			initialCarsInDatabase[0].Vin,
			0,
			"Car with VIN {JH4DB1561NS000567} is not available for the specified time range",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rental, err := customerOperations.RentCar("customerID", tc.start, tc.end, tc.vin)

			// No error expected
			if tc.expectedErrorMessage == "" {

				if err != nil {
					t.Fatalf("Test case %v should not create error, but creates error \"%s\"", tc.name, err.Error())
				}

				// check if the price was calculated correctly
				if rental.Price != tc.expectedPrice {
					t.Fatalf("Test case %v should calculate price %f but calculates price %f", tc.name, tc.expectedPrice, rental.Price)
				}
			}

			// Error mismatch
			if tc.expectedErrorMessage != "" && (err == nil || err.Error() != tc.expectedErrorMessage) {
				t.Fatalf("Test case %v should produce error with message: %v, but instead produces: %v", tc.name, tc.expectedErrorMessage, err.Error())
			}
		})
	}
}
