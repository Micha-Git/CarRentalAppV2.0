# Use Case "Add Car to Fleet"

```
Title: Add Car to Fleet

Primary Actors: Fleet manager
Secondary Actors: None

Preconditions:
    - Fleet manager is assigned to the fleet
Postconditions:
    - The vin, model, and brand of the car are stored.
    - The car is added to the fleet.
    - The car is available for rental.

Flow:
1. Fleet manager adds a new car to the Fleet by providing a VIN.
2. System gathers the basic car information from the external connected car system.
3. System stores the basic car information.
4. Fleet manager is informed about the success of the operation.

Alternative flows:
1a. Fleet manager provides an invalid VIN.
    1a1. System informs the fleet manager about the invalid VIN.
    1a2. The use case is terminated.

Information Requirements:
    Connected Cars System:
    - brand
    - model
```
