import { Component, Input, OnInit } from '@angular/core';

import { Song } from '../../core';

@Component({
  selector: 'app-songs-preview',
  templateUrl: './songs-preview.component.html'
})
export class SongsPreviewComponent implements OnInit{
  @Input() song: Song;
  ngOnInit() {
  }
}
