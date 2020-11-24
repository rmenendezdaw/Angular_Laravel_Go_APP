import { Component, Input } from '@angular/core';

import { Song } from '../../core';

@Component({
  selector: 'app-songs-preview',
  templateUrl: './songs-preview.component.html'
})
export class SongPreviewComponent {
  @Input() song: Song;

  onToggleFavorite(favorited: boolean) {
    this.song['favorited'] = favorited;

    if (favorited) {
      this.song['favoritesCount']++;
    } else {
      this.song['favoritesCount']--;
    }
  }
}
