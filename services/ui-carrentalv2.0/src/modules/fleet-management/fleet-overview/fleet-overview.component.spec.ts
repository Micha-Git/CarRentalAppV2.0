import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FleetOverviewComponent } from './fleet-overview.component';

describe('FleetManagementComponent', () => {
  let component: FleetOverviewComponent;
  let fixture: ComponentFixture<FleetOverviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FleetOverviewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(FleetOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
