# Class diagram UI Fleet Management 

![](figures/cld-fleet-management-services.png)

The figure shows all the Angular specific elements used in the context of Fleet Management.

(\<<component\>> NavBarComponent) The NavBarComponent controls the tabs, the user is able to access. Therefore it is closely linked to FleetManagement classes

(\<<component\>>FleetOverviewComponent) The FleetOverviewComponent is responsible for the listing of all cars in the fleet. The use cases "List Cars In Fleet", "View Car Information" and "Remove Car from Fleet" will be implemented here.

(\<<component\>> CarAdditionComponent) This component shows an entry field for the user to enter information for a new car.


(\<<component\>> FleetManagementAPIService) This Service is created using the protocol buffer file. This services provides a client which allows other components to use the methods defined in the protocol buffer file, without any other difficulites.