import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {
    CustomerServiceClient, RentableCarsCollectionServiceClient,
    RentalsCollectionServiceClient
} from "./proto/Api_specification_am_rental_managementServiceClientPb";
import {environment} from "../../../environments/environment";
import {
    AddCarToRentalRequest,
    CancelRentalRequest,
    ListAvailableCarsRequest,
    ListCarRentalsRequest, ListCustomerRentalsRequest,
    RemoveRentableCarRequest,
    RentableCar,
    Rental,
    RentCarRequest,
    Vin
} from "./proto/api_specification_am_rental_management_pb";
import {handleServiceResponse} from "./response-handler";
import {Injectable} from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class RentalManagementService {
    private readonly customerServiceClient: CustomerServiceClient;
    private readonly rentalsCollectionServiceClient: RentalsCollectionServiceClient;
    private readonly rentableCarsCollectionServiceClient: RentableCarsCollectionServiceClient;

    constructor() {
        this.customerServiceClient = new CustomerServiceClient(environment.rentalManagementServiceHost);
        this.rentalsCollectionServiceClient = new RentalsCollectionServiceClient(environment.rentalManagementServiceHost);
        this.rentableCarsCollectionServiceClient = new RentableCarsCollectionServiceClient(environment.rentalManagementServiceHost);
    }

    /* Actions provided by customerServiceClient */
    rentCar(customerId: string, startDate: Timestamp, endDate: Timestamp, vin: Vin): Promise<Rental> {
        const request = new RentCarRequest();
        request.setCustomerid(customerId);
        request.setStartdate(startDate);
        request.setEnddate(endDate);
        request.setVin(vin);

        return new Promise<Rental>((resolve, reject) => {
            this.customerServiceClient.rentCar(request, null, (err, response) => {
                handleServiceResponse<Rental>('renting a car', err, response,
                    response?.getError(), response?.getRental(), resolve, reject);
            })
        });
    }

    cancelRental(customerId: string, rentalId: string): Promise<boolean> {
        const request = new CancelRentalRequest();
        request.setCustomerid(customerId);
        request.setRentalid(rentalId);

        return new Promise<boolean>((resolve, reject) => {
            this.customerServiceClient.cancelRental(request, null, (err, response) => {
                handleServiceResponse<boolean>('canceling a rental', err, response,
                    response?.getError(), response?.getResult(), resolve, reject);
            })
        });
    }

    /* Actions provided by rentalsCollectionServiceClient */
    listAvailableCars(startDate: Timestamp, endDate: Timestamp, location: string): Promise<RentableCar[]> {
        const request = new ListAvailableCarsRequest();
        request.setStartdate(startDate);
        request.setEnddate(endDate);
        request.setLocation(location);

        return new Promise<RentableCar[]>((resolve, reject) => {
            this.rentalsCollectionServiceClient.listAvailableCars(request, null, (err, response) => {
                handleServiceResponse<RentableCar[]>("listing available cars", err, response,
                    response?.getError(), response?.getCarsList(), resolve, reject);
            });
        });
    }

    listCarRentals(vin: Vin): Promise<Rental[]> {
        const request = new ListCarRentalsRequest();
        request.setVin(vin);

        return new Promise<Rental[]>((resolve, reject) => {
            this.rentalsCollectionServiceClient.listCarRentals(request, null, (err, response) => {
                handleServiceResponse<Rental[]>("listing car rentals", err, response,
                    response?.getError(), response?.getRentalsList(), resolve, reject);
            });
        });
    }

    listCustomerRentals(customerId: string): Promise<Rental[]> {
        const request = new ListCustomerRentalsRequest();
        request.setCustomerid(customerId);

        return new Promise<Rental[]>((resolve, reject) => {
            this.rentalsCollectionServiceClient.listCustomerRentals(request, null, (err, response) => {
                handleServiceResponse<Rental[]>("listing customer rentals", err, response,
                    response?.getError(), response?.getRentalsList(), resolve, reject);
            });
        });
    }

    /* Actions provided by rentableCarsCollectionServiceClient */
    addCarToRental(vin: Vin, location: string): Promise<RentableCar> {
        const request = new AddCarToRentalRequest();
        request.setVin(vin)
        request.setLocation(location);

        return new Promise<RentableCar>((resolve, reject) => {
            this.rentableCarsCollectionServiceClient.addCarToRental(request, null, (err, response) => {
                handleServiceResponse<RentableCar>('adding car to rental', err, response,
                    response?.getError(), response?.getCar(), resolve, reject);
            })
        });
    }

    removeRentableCar(vin: Vin): Promise<boolean> {
        const request = new RemoveRentableCarRequest();
        request.setVin(vin);


        return new Promise<boolean>((resolve, reject) => {
            this.rentableCarsCollectionServiceClient.removeRentableCar(request, null, (err, response) => {
                handleServiceResponse<boolean>('adding car to rental', err, response,
                    response?.getError(), response?.getResult(), resolve, reject);
            })
        });
    }

}