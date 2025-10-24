import {Component} from '@angular/core';


@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.scss']
})
export class LandingPageComponent {

  constructor() {
    localStorage.setItem('navBarVisibility', 'false');
  }

  showFleetManagerView() {
    localStorage.setItem('navBarVisibility', 'true');
    localStorage.setItem('user', 'fleetManager');

  }

  showCustomerView() {
    localStorage.setItem('navBarVisibility', 'true');
    localStorage.setItem('user', 'customer');
  }

  visibilityOfNavBar(): string {
    return localStorage.getItem('navBarVisibility')!;
  }
}