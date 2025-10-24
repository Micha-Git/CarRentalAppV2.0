import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AvailableCarsListComponent } from './available-cars-list.component';

describe('AvailableCarsListComponent', () => {
  let component: AvailableCarsListComponent;
  let fixture: ComponentFixture<AvailableCarsListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AvailableCarsListComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(AvailableCarsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
