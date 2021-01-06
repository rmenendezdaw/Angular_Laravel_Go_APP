import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Song, SongsService } from '../../../core';

@Component({
  selector: 'app-songs-admin-preview',
  templateUrl: './songs-admin-preview.html'
})
export class SongsAdminPreviewComponent {
  constructor (
    private songsService: SongsService,
    private router: Router

  ) {}

  @Input() song: Song;
  deleteSong() {

    this.songsService.destroy(this.song.id)
      .subscribe(
        success => {
          console.log('Done');
          this.router.navigateByUrl('/admin');
        }
      );
  }
}
