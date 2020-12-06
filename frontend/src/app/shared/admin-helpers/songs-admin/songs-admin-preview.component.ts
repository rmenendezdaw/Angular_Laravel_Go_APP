import { Component, Input, OnInit } from '@angular/core';

import { Song } from '../../../core';

@Component({
  selector: 'app-songs-admin-preview',
  templateUrl: './songs-admin-preview.html'
})
export class SongsAdminPreviewComponent implements OnInit{
  @Input() song: Song;
  ngOnInit() {
        
  }
}
