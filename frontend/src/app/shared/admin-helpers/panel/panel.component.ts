import { Component, OnInit, Input } from '@angular/core';
import { PanelService } from '../../../core';

@Component({
  selector: 'app-panel',
  templateUrl: './panel.component.html'
})
export class PanelComponent implements OnInit {
  constructor(
    private panelService: PanelService) { }

  users: number;

  ngOnInit() {
    this.panelService.get().subscribe(data => {
      console.log(data);
      this.users = data;
    })
  }
}
