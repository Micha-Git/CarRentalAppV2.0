import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RentalOverviewComponent } from './rental-overview.component';

describe('RentalOverviewComponent', () => {
  let component: RentalOverviewComponent;
  let fixture: ComponentFixture<RentalOverviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RentalOverviewComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(RentalOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
