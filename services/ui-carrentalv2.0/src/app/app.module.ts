import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {NavBarModule} from '../modules/nav-bar/nav-bar.module';
import {CommonModule} from '@angular/common';
import {AppComponent} from "./app.component";
import {LandingPageModule} from "../modules/landing-page/landing-page.module";
import {AppRoutingModule} from "./app-routing.module";

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    NavBarModule,
    AppRoutingModule,
    LandingPageModule
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
