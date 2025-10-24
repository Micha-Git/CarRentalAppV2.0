import {Component, OnInit} from '@angular/core';
import {DatePipe, NgForOf, NgIf} from '@angular/common';
import {ActivatedRoute} from '@angular/router';
import {FleetManagementService} from '../../fleet-management/service/fleet-management.service';
import {RentalManagementService} from '../service/rental-management.service';
import {Rental} from '../service/proto/api_specification_am_rental_management_pb';
import {Vin} from '../../fleet-management/service/proto/api_specification_am_fleet_management_pb';

@Component({
  selector: 'app-list-car-rentals',
  standalone: true,
  imports: [NgForOf, NgIf, DatePipe],
  templateUrl: './list-car-rentals.component.html',
  styleUrls: ['./list-car-rentals.component.scss'],
  providers: [FleetManagementService],
})
export class ListCarRentalsComponent implements OnInit {
  carRentals: VisualRental[] = [];
  vin: Vin | null = null;

  constructor(private rentalManagementService: RentalManagementService, private route: ActivatedRoute) {
  }

  ngOnInit() {
    const vinString = this.route.snapshot.queryParamMap.get('vin');
    this.vin = new Vin();
    this.vin.setVin(vinString!);

    if (this.vin) {
      this.listCarRentals(this.vin);
    }
  }

  protected formatDate(date: Date): string {
    let day = '' + date.getDate();
    let month = '' + (date.getMonth() + 1);
    let year = '' + date.getFullYear();

    day = day.length < 2 ? `0${day}` : day;
    month = month.length < 2 ? `0${month}` : month;
    return day + '.' + month + '.' + year;
  }

  private listCarRentals(vin: Vin) {
    this.rentalManagementService
      .listCarRentals(vin)
      .then((rentals: Rental[]) => {
        this.carRentals = [];
        for (const rental of rentals) {
          this.carRentals.push(this.mapRentalToVisual(rental));
        }
      })
      .catch((error: any) => console.error('Error listing car rentals:', error));
  }

  private mapRentalToVisual(rental: Rental): VisualRental {
    return {
      startDate: rental.getStartdate()!.toDate(),
      endDate: rental.getEnddate()!.toDate(),
      customer: rental.getCustomerid()!,
    };
  }
}

interface VisualRental {
  startDate: Date;
  endDate: Date;
  customer: string;
}
