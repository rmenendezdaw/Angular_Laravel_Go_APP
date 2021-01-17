import { Component, OnInit } from '@angular/core';
import { User, UserService } from '../../../core';


@Component({
  selector: 'app-users-admin',
  templateUrl: './users-admin.component.html',
  styleUrls: ['./users-admin.component.css']
})
export class UsersAdminComponent implements OnInit {

  constructor(
    private userService: UserService
  ) { }

  results: User[];

  ngOnInit() {
    this.results = [];
    this.userService.getAllUsers().subscribe(data => {
      this.results = data;
    })
  }

}
