import {ComponentFixture, TestBed} from '@angular/core/testing';
import {ListCarRentalsComponent} from "./list-car-rentals.component";


describe('ListCarRentalsComponent', () => {
  let component: ListCarRentalsComponent;
  let fixture: ComponentFixture<ListCarRentalsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListCarRentalsComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(ListCarRentalsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
