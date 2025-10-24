import { Injectable } from '@angular/core';
import { AddCarToFleetRequest, Car, ErrorDetail, ListCarsInFleetRequest, RemoveCarFromFleetRequest, ViewCarInformationRequest, Vin } from './proto/api_specification_am_fleet_management_pb';
import { handleServiceResponse } from './response-handler';
import { RpcError } from 'grpc-web';
import { environment } from '../../../environments/environment';
import { CarServiceClient, FleetServiceClient } from './proto/Api_specification_am_fleet_managementServiceClientPb';

@Injectable({
  providedIn: 'root'
})
export class FleetManagementService {

  private readonly carServiceClient: CarServiceClient;
  private readonly fleetServiceClient: FleetServiceClient;

  constructor() {
      this.carServiceClient = new CarServiceClient(environment.fleetManagementServiceHost);
      this.fleetServiceClient = new FleetServiceClient(environment.fleetManagementServiceHost);
  }


  public listCarsInFleet(): Promise<Car[]> {

    const req = new ListCarsInFleetRequest();
    //Change to fleet of fleetManager who is logged In
    req.setFleetid("1");
    const fleetCars: Car[] = [];
    
    return new Promise<Car[]>((resolve, reject) => {
            this.fleetServiceClient.listCarsInFleet(req, null, (err, response) => {
                handleServiceResponse<Car[]>("listing cars in fleet", err, response,
                    response?.getError(), response?.getCarsList(), resolve, reject);   
                    console.log(response);               
                });
        });
  }



  public addCarToFleet(location: string, vinInput: string, fleetId: string): Promise<Car> {
  
    const req = new AddCarToFleetRequest();
    const vin = new Vin();
    vin.setVin(vinInput);
    req.setVin(vin);
    req.setFleetid(fleetId);
    req.setLocation(location);


    return new Promise<Car>((resolve, reject) => {
      
      this.fleetServiceClient.addCarToFleet(req, null, (err, response) => {
          handleServiceResponse<Car>("add car to fleet", err, response,
              response?.getError(), response?.getCar(), resolve, reject);
              window.location.reload();
      });
  });
  }

  public removeCarFromFleet(vinInput: string): Promise<boolean>{
    const req = new RemoveCarFromFleetRequest();
    const vin = new Vin();
    vin.setVin(vinInput);
    req.setVin(vin);

    return new Promise<boolean>((resolve, reject) => {
      this.fleetServiceClient.removeCarFromFleet(req, null, (err, response) => {
          window.location.reload();
      });
  });

  }
}
