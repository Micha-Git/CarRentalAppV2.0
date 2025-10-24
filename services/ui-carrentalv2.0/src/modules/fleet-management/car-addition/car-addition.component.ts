import { Component } from '@angular/core';
//import { FleetManagementService } from '../service/fleet-management.service';
import { NgForOf, NgIf } from '@angular/common';
import { RouterLink } from '@angular/router';
import { FleetManagementService } from '../service/fleet-management.service';

@Component({
  selector: 'app-car-addition',
  standalone: true,
  imports: [
    NgForOf,
    NgIf,
    RouterLink
  ],
  providers: [FleetManagementService],

  templateUrl: './car-addition.component.html',
  styleUrl: './car-addition.component.scss'
})
export class CarAdditionComponent {


  constructor(private fleetService: FleetManagementService){}

  addCarToFleet(location: string, vin: string, fleetId: string) {
    this.fleetService.addCarToFleet(location, vin, fleetId).then(window.location.reload);
  }

}
