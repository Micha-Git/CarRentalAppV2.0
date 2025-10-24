import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LandingPageComponent} from '../modules/landing-page/landing-page/landing-page.component';
import {RentalOverviewComponent} from "../modules/rental-management/rental-overview/rental-overview.component";
import { AvailableCarsListComponent } from '../modules/rental-management/available-cars-list/available-cars-list.component';
import {RentalCreationComponent} from "../modules/rental-management/rental-creation/rental-creation.component";
import { FleetOverviewComponent } from '../modules/fleet-management/fleet-overview/fleet-overview.component';
import { CarAdditionComponent } from '../modules/fleet-management/car-addition/car-addition.component';
import {ListCarRentalsComponent} from "../modules/rental-management/list-car-rentals/list-car-rentals.component";

const routes: Routes = [
  // Home Path
  {path: '', redirectTo: 'landing-page', pathMatch: 'full'},

  // Path to landing page
  {
    path: 'landing-page',
    component: LandingPageComponent,
    loadChildren: () => import('../modules/landing-page/landing-page.module').then(m => m.LandingPageModule)
  },

  // Path to rental management
  //{
  // path: 'service',
  //  component:
  //},

  {
    path: 'rental-overview',
    component: RentalOverviewComponent
  },

  {
    path: 'rental-creation',
    component: RentalCreationComponent
  },

  {
    path: 'available-cars-list',
    component: AvailableCarsListComponent
  },

  // Path to list car rentals
  {
    path: 'list-car-rentals',
    component: ListCarRentalsComponent
  },

  // Path to fleet management
  //{
  //  path: 'fleet-management',
  //  component:
  //}
  // Path to fleet overview
  {
    path: 'fleet-overview',
    component: FleetOverviewComponent
  }, 

  // Path to car addition
  {
    path: 'car-addition',
    component: CarAdditionComponent

  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
