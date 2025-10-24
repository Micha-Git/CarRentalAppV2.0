# CarRentalAppV2.0

CarRentalApp is a microservice-based application which enables to rent a car from a rental company. In comnparison to its earlier version, CarRentalAppV2.0 is extended by an application microservice AM-FleetManagement by which the cars to be rented are organized in fleets.

## ANALYSIS

[Terms Used V2.0](pages/terms_used_v2.0.md)  
[Story V2.0](pages/story_v2.0.md)  
[Capability "Management of Rentals" V2.0](./pages/cap_management_of_rentals_v2.0.md)  
[Capability "Management of Fleets" V2.0](./pages/cap_management_of_fleets_v2.0.md)  

[Use Case Diagram V2.0](pages/use_case_diagram_v2.0.md)  

[Use Case "List Available Cars" V2.0](./pages/uc_list_available_cars_v2.0.md)  
[Use Case "Rent a Car" V2.0](./pages/uc_rent_a_car_v2.0.md)  
[Use Case "List Customer Rentals" V2.0](./pages/uc_list_customer_rentals_v2.0.md)

[Use Case "List Car Rentals" V2.0](./pages/uc_list_car_rentals_v2.0.md)  
[Use Case "Add Car to Fleet" V2.0](./pages/uc_add_car_to_fleet_v2.0.md)  
[Use Case "Remove Car from Fleet" V2.0](./pages/uc_remove_car_from_fleet_v2.0.md)  

## DESIGN

[Component Diagram CarRentalAppV2.0](./pages/cd_car_rental_app_v2.0.md)  
[Component Diagram UI-CarRentalV2.0](./pages/cd_ui-car_rental_v2.0.md)  

[API Diagram DM-CarV2.0](./pages/ad_dm-car_v2.0.md)  
[API Specification DM-CarV2.0](https://github.com/Micha-Git/CarRentalAppV2.0/blob/main/services/dm-carv2.0/src/api/specification/openapi.yaml)

[API Diagram AM-RentalManagementV2.0](./pages/ad_am-rental_management_v2.0.md)  
[API Specification AM-RentalManagementV2.0](https://github.com/Micha-Git/CarRentalAppV2.0/blob/main/services/am-rentalmanagementv2.0/src/api/specification/api_specification_am_rental_management.proto)  

[API Diagram AM-FleetManagementV1.0](./pages/ad_am-fleet_management_v1.0.md)  
[API Specification AM-FleetManagementV1.0](https://github.com/Micha-Git/CarRentalAppV2.0/blob/main/services/am-fleetmanagementv1.0/src/api/specification/api_specification_am_fleet_management.proto?ref_type=heads)  

[Extended Component Diagram AddCarToFleetV2.0](./pages/ecd_add_car_to_fleet_v2.0.md)  
[Extended Component Diagram RemoveCarFromFleetV2.0](./pages/ecd_remove_car_from_fleet_v2.0.md)  

[Orchestration Diagram AddCartoFleetV2.0](./pages/od_add_car_to_fleet_v2.0.md)  
[Orchestration Diagram RemoveCarfromFleetV2.0](./pages/od_remove_car_from_fleet_v2.0.md)  

## Implementation

[Microservice Implementation DM-CarV2.0](https://github.com/Micha-Git/CarRentalAppV2.0/tree/main/services/dm-carv2.0)

[Microservice Implementation AM-RentalManagementV2.0](https://github.com/Micha-Git/CarRentalAppV2.0/tree/main/services/am-rentalmanagementv2.0)  

[Microservice Implementation AM-FleetManagementV1.0](https://github.com/Micha-Git/CarRentalAppV2.0/tree/main/services/am-fleetmanagementv1.0)  

[User Interface UI-CarRentalV2.0](https://github.com/Micha-Git/CarRentalAppV2.0/tree/main/services/ui-carrentalv2.0)  

## DEPLOYMENT AND OPERATIONS

[SPS Diagram CarRentalApp V2.0](./pages/sps_diagram_car_rental_app_v2.0.md)  