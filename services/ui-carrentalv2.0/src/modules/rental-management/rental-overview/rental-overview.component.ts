import {Component} from '@angular/core';
import {DatePipe, NgForOf, NgIf} from "@angular/common";
import {RouterLink} from "@angular/router";
import {RentalManagementService} from "../service/rental-management.service";
import {Rental} from "../service/proto/api_specification_am_rental_management_pb";

@Component({
  selector: 'app-rental-overview',
  standalone: true,
  imports: [
    NgForOf,
    NgIf,
    RouterLink
  ],
  providers: [DatePipe],
  templateUrl: './rental-overview.component.html',
  styleUrl: './rental-overview.component.scss'
})
export class RentalOverviewComponent {

  rentals: VisualRental[] = [];

  constructor(private rentalManagementService: RentalManagementService) {}

  ngOnInit() {
    this.updateRentals()
  }

  cancelRental(rental: VisualRental) {
    this.rentalManagementService.cancelRental("CUSTOMER_ID", rental.id)
        .then(() => {
          this.updateRentals()
        });
  }

  protected formatDate(date: Date): string {
    let day = "" + date.getDate();
    let month = "" + (date.getMonth() + 1);
    let year = "" + date.getFullYear();

    day = day.length < 2 ? `0${day}` : day;
    month = month.length < 2 ? `0${month}` : month;
    return day + "." + month + "." + year;
  }

  private updateRentals() {
    this.rentalManagementService.listCustomerRentals("CUSTOMER_ID")
        .then((customerRentals: Rental[]) => {

          this.rentals = [];
          for (const rental of customerRentals) {
            this.rentals.push(this.mapRentalToVisual(rental))
          }
    });
  }

  private mapRentalToVisual(rental: Rental): VisualRental {
    return {
      id: rental.getId(),
      vin: rental.getCar()!.getVin()!.getVin(),
      brand: rental.getCar()!.getBrand(),
      model: rental.getCar()!.getModel(),
      logoUrl: "assets/dummycar.png",
      price: rental.getPrice(),
      rentalStart: rental.getStartdate()!.toDate(),
      rentalEnd: rental.getEnddate()!.toDate()
    };
  }

  protected readonly localStorage = localStorage;
}

interface VisualRental {
  id: string,
  vin: string,
  brand: string,
  model: string,
  logoUrl: string,
  price: number,
  rentalStart: Date,
  rentalEnd: Date
}
