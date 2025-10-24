import {Component} from '@angular/core';
import {FleetManagementService} from '../service/fleet-management.service';
import {Car} from '../service/proto/api_specification_am_fleet_management_pb';
import {NgForOf, NgIf} from '@angular/common';
import {Router, RouterLink} from '@angular/router';

@Component({
  selector: 'app-fleet-overview',
  standalone: true,
  imports: [NgForOf,
    NgIf,
    RouterLink
  ],
  templateUrl: './fleet-overview.component.html',
  styleUrl: './fleet-overview.component.scss',
  providers: [FleetManagementService]
})
export class FleetOverviewComponent {

  fleetCars: VisualCar[] = [];

  constructor(private fleetService: FleetManagementService, private router: Router) {
  }

  ngOnInit() {
    this.listCars();
  }

  showAddCar() {
    window.location.href = "car-addition"
  }

  removeCar(vin: string) {
    this.fleetService.removeCarFromFleet(vin).then(this.listCars);
  }

  viewCarRentals(vin: string) {
    this.router.navigate(['/list-car-rentals'], {queryParams: {vin: vin}});
  }

  private listCars() {

    this.fleetService
      .listCarsInFleet()
      .then((loadedCars: Car[]) => {
        this.fleetCars = [];
        for (const fleetCar of loadedCars) {
          this.fleetCars.push(this.mapCarToVisual(fleetCar))
        }
      });
  }

  private mapCarToVisual(fleetCar: Car): VisualCar {
    return {
      vin: fleetCar.getVin()!.getVin(),
      brand: fleetCar.getBrand(),
      model: fleetCar.getModel(),
    };
  }
}

interface VisualCar {
  vin: string,
  brand: string
  model: string,
}