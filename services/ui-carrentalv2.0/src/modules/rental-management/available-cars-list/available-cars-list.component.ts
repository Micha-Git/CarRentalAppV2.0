import {Component} from '@angular/core';
import {NgForOf, NgIf} from "@angular/common";
import {Router, RouterLink} from '@angular/router';
import {RentalManagementService} from "../service/rental-management.service";
import {RentableCar, Vin} from "../service/proto/api_specification_am_rental_management_pb";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";

@Component({
  selector: 'app-available-cars-list',
  standalone: true,
  imports: [
    NgForOf,
    NgIf,
    RouterLink
  ],
  templateUrl: './available-cars-list.component.html',
  styleUrl: './available-cars-list.component.scss'
})
export class AvailableCarsListComponent {

  searchFilters: AvailableCarsSearchFilter = {
    availableFrom: null,
    availableTo: null,
    location: "Karlsruhe"
  }

  availableCars: VisualRentableCar[] = [];

  constructor(private router: Router, private rentalManagementService: RentalManagementService) {
  }

  rentalCar(car: VisualRentableCar) {
    if (!this.searchFilters.location || !this.searchFilters.availableFrom || !this.searchFilters.availableTo) {
      return;
    }

    const vin: Vin = new Vin()
    vin.setVin(car.vin)

    this.rentalManagementService
      .rentCar("CUSTOMER_ID", this.searchFilters.availableFrom, this.searchFilters.availableTo, vin)
      .then(() => {
        this.updateAvailableCars();
      });
  }

  onRentalCar(car: VisualRentableCar) {
    this.router.navigate(['/rental-creation'], {
      queryParams: {
        availableFrom: this.searchFilters.availableFrom!.toDate().toISOString(),
        availableTo: this.searchFilters.availableTo!.toDate().toISOString(),
        location: location,
        car: JSON.stringify(car),
      }
    });
  }


  onSearchCars(availableFrom: Date | null, availableTo: Date | null, location: string) {

    if (!availableFrom || !availableTo) {
      return;
    }

    this.searchFilters = {
      availableFrom: Timestamp.fromDate(availableFrom),
      availableTo: Timestamp.fromDate(availableTo),
      location: location
    }

    this.updateAvailableCars();
  }

  protected formatTimestamp(timestamp: Timestamp): string {
    const date = timestamp.toDate()

    let day = "" + date.getDate();
    let month = "" + (date.getMonth() + 1);
    let year = "" + date.getFullYear();

    day = day.length < 2 ? `0${day}` : day;
    month = month.length < 2 ? `0${month}` : month;
    return day + "." + month + "." + year;
  }

  private updateAvailableCars() {
    this.rentalManagementService
      .listAvailableCars(this.searchFilters.availableFrom!, this.searchFilters.availableTo!, this.searchFilters.location!)
      .then((rentableCars: RentableCar[]) => {

        this.availableCars = [];
        for (const rentableCar of rentableCars) {
          this.availableCars.push(this.mapRentableCarToVisual(rentableCar))
        }
      });
  }

  private mapRentableCarToVisual(rentableCar: RentableCar): VisualRentableCar {
    return {
      vin: rentableCar.getVin()!.getVin(),
      brand: rentableCar.getBrand(),
      model: rentableCar.getModel(),
      logoUrl: "assets/dummycar.png",
      rentalPricePerDay: rentableCar.getPriceperday(),
    };
  }

  changePeriod() {
    this.searchFilters = {
      availableFrom: null,
      availableTo: null,
      location: null
    };
  }

  protected readonly localStorage = localStorage;
}

interface AvailableCarsSearchFilter {
  availableFrom: Timestamp | null,
  availableTo: Timestamp | null,
  location: string | null
}

interface VisualRentableCar {
  vin: string,
  brand: string,
  model: string,
  logoUrl: string,
  rentalPricePerDay: number
}
