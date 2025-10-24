import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {NavBarComponent} from './nav-bar/nav-bar.component';
import {RouterLink} from "@angular/router";



@NgModule({
  imports: [
    CommonModule,
    RouterLink,
  ],
  declarations: [NavBarComponent],
  exports: [NavBarComponent]
})


export class NavBarModule {
}

