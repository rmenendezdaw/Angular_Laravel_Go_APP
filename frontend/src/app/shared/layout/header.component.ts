import { Component, OnInit } from '@angular/core';

import { User, UserService } from '../../core';

@Component({
  selector: 'app-layout-header',
  templateUrl: './header.component.html'
})
export class HeaderComponent implements OnInit {
  constructor(
    private userService: UserService
  ) {}

  currentUser: User;
  type: String;

  ngOnInit() {
    this.userService.currentUser.subscribe(
      (userData) => {
        console.log("TYPEEEEEEE")
        console.log(userData)
        this.currentUser = userData;
        this.type = userData.type;
      }
    );
  }
}
