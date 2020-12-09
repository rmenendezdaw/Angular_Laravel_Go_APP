import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Song, SongsService } from '../core';


@Component({
  selector: 'app-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {
  song: Song;
  isDeleting = false;

  constructor(
    private songsService: SongsService,
    private router: Router
  ) { }

  ngOnInit() {
    console.log('hi');
  }
  deleteSong() {
    this.isDeleting = true;

    this.songsService.destroy(this.song.id)
      .subscribe(
        success => {
          this.router.navigateByUrl('/');
        }
      );
  }
}
