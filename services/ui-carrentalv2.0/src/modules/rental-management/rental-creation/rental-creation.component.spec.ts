import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RentalCreationComponent } from './rental-creation.component';

describe('RentalCreationComponent', () => {
  let component: RentalCreationComponent;
  let fixture: ComponentFixture<RentalCreationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RentalCreationComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(RentalCreationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
