import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Song, SongsService } from '../../../core';

@Component({
  selector: 'app-songs-admin-preview',
  styleUrls: ['./songs-admin.component.css'],
  templateUrl: './songs-admin-preview.html'
})
export class SongsAdminPreviewComponent {
  constructor (
    private songsService: SongsService,
    private router: Router
  ) {}

  @Input() song: Song;

  deleteSong(event) {
    event.stopPropagation()

    this.songsService.destroy(this.song.id)
      .subscribe(
        success => {
          this.router.navigateByUrl('.', { skipLocationChange: true })
          .then(() => {
            this.router.navigate(['/']);
        }); 
        }
      );
  }

  modifySong(event, id) {
    event.stopPropagation()
    this.router.navigateByUrl("/editor/" + id);
  }
}
