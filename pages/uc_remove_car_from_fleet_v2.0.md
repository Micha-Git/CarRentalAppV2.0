# Use Case "Remove Car from Fleet"

```
Title: Remove Car from Fleet

Primary Actors: Fleet manager
Secondary Actors: None

Preconditions:
    - Fleet for the fleet manager exists
    - Car is in the fleet
Postconditions:
    - The vin, model, and brand of the car are removed from the system.
    - The car is removed from the fleet.
    - The car is not available for rental.

Flow:
1. Fleet manager lists all cars in the fleet by using the use case "List Cars in Fleet" and selects a car.
2. System removes the car from the fleet 
3. System informs the fleet manager about the success of the operation.

Alternative flows:
2a. Rentals exist for the car
    2a1. System informs the fleet manager about the existing rentals.
    2a2. The use case is terminated.

Information Requirements: None
```
