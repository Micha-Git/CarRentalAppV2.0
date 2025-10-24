# Code Sketch AM-RentalManagementV2.0 API
![](../figures/fs_am_rental_management_v2.0_api.png)

The package api contains the controllers representing the services defined in the api specification,
two generated protobuf files containing functions and models to interact with gRPC clients, 
and mappers to convert between the protobuf entities and the models they represent. 
The Controllers are illustrated in the following figures and their descriptions.

## Customer Controller
![](../figures/cs_customer_collection_controller.png)

The CustomerController uses the CustomerOperationsInterface to provided its functions:
RentCar, and CancelRental.

## Rentals Collection Controller
![](../figures/cs_rentals_collection_controller.png)

The RentalsCollectionController uses the RentalsCollectionOperationsInterface to provided its functions:
ListAvailableCars, ListCarRentals, and ListCustomerRentals.

## Rentable Cars Collection Controller
![](../figures/cs_rentable_cars_collection_controller.png)

The RentableCarsCollectionController uses the RentableCarsCollectionOperationsInterface to provided its functions: 
AddCarToRental, and RemoveRentableCar.