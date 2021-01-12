import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '../../core';

@Component({
  selector: 'app-layout-header',
  templateUrl: './header.component.html'
})
export class HeaderComponent implements OnInit {
  constructor(
    private userService: UserService,
    private router: Router,

  ) {}

  currentUser: User;
  type: String;

  ngOnInit() {
    this.userService.currentUser.subscribe(
      (userData) => {
        let route = (userData.type === 'admin') ? '/admin' : '/';

        this.currentUser = userData;
        this.type = userData.type;
        //sthis.router.navigateByUrl(route)
      }
    );
  }
}
