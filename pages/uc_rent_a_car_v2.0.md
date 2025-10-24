# Use Case "Rent a Car"

```
Title: Rent a Car

Primary Actors: Customer
Secondary Actors: None

Preconditions:
    - System has cars that are registered to a fleet in the selected location and available in the selected rental period
Postconditions:
    - A rental is created

Flow:
1. Customer lists all available cars using the use case "List Avaibale Cars" and selects a car.
2. System checks if the car is available for the given time slot.
3. System creates a new rental and returns it.

Alternative flows:
2a. Given an invalid time
    2a1. System prompts the customer to select another time
2b. Given an invalid vin
    2b1. System prompts customer to select another car



Information Requirements: None
```
