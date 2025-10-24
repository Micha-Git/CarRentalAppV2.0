# Use Case "List Car Rentals"

```
Title: List Car Rentals

Primary Actors: Fleet manager
Secondary Actors: None

Preconditions:
    - Fleet for fleet manager exists
    - The car exists in the fleet
Postconditions: None

Flow:
1. Fleet manager selects to view all rentals of the car.
2. System presents all rentals of the car including the start date, end date, and the customer.

Alternative flows:
2a. The car has no rentals.
    1. System informs the fleet manager that there are no rentals.

Information Requirements: None
```
