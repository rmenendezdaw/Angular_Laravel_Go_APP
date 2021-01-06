import { Component, OnInit } from '@angular/core';
import { UserService } from '../core';

@Component({
  selector: 'app-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {

  constructor(
    private userService: UserService
  ) { }

  ngOnInit() {
    console.log('panel component');
    this.userService.getUser()
  }
}
