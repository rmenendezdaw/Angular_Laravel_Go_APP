import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SongsAdminComponent } from './songs-admin.component';

describe('SongsAdminComponent', () => {
  let component: SongsAdminComponent;
  let fixture: ComponentFixture<SongsAdminComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SongsAdminComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SongsAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
