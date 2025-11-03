# Use Case "Cancel Rental"

```
Title: Cancel Rental

Primary Actors: Customer
Secondary Actors: None

Preconditions:
    - A rental with the wanted car was booked by the customer.
    
Postconditions:
    - The rental is canceled

Flow:
1. Customer lists all rentals using the use case "List my Car Rentals" and selects a rental.
2. System asks whether the user is sure to cancel his rental.
3. System cancels the rental.

Alternative flows:
1a. Rental was already canceled before.
    1. System tells the user that the rental is already canceled.

Information Requirements: None
```