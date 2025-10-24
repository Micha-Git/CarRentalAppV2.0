package operations

import (
	"fmt"
	"testing"
	"time"
)

func TestCancelRental(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	// Rent a car
	location := time.Now().Location()
	startDate := time.Date(2024, 1, 1, 8, 0, 0, 0, location)
	endDate := time.Date(2024, 1, 1, 22, 0, 0, 0, location)
	storedRental, err := customerOperations.RentCar("0", startDate, endDate, initialCarsInDatabase[0].Vin) // Tesla Model 3 Mannheim
	if err != nil {
		t.Fatalf("Failed to rent car with vin '%v': %v", initialCarsInDatabase[0].Vin, err)
	}

	cases := []struct {
		name                 string
		customerId           string
		rentalId             string
		expectedErrorMessage string
	}{
		{
			"NonExistentRental",
			"invalidId",
			"invalidId",
			fmt.Sprintf(
				"Failed to cancel the rental with ID %s: Database failed to delete rental with ID %s belonging to customer with ID %s: There is no rental matching the given condition",
				"invalidId", "invalidId", "invalidId"),
		},
		{
			"RentalBelongsToDifferentCustomer",
			"wrongCustomer",
			storedRental.Id,
			fmt.Sprintf(
				"Failed to cancel the rental with ID %s: Database failed to delete rental with ID %s belonging to customer with ID %s: There is no rental matching the given condition",
				storedRental.Id, storedRental.Id, "wrongCustomer"),
		},
		{
			"success",
			storedRental.CustomerId,
			storedRental.Id,
			"",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := customerOperations.CancelRental(tc.customerId, tc.rentalId)

			// No error expected
			if tc.expectedErrorMessage == "" && result != nil {
				t.Fatalf("Test case %v should not create error", tc.name)
			}

			// Error expected
			if tc.expectedErrorMessage != "" && (result == nil || result.Error() != tc.expectedErrorMessage) {
				t.Fatalf("Test case %v should produce error with message: %v", tc.name, tc.expectedErrorMessage)
			}
		})
	}
}
