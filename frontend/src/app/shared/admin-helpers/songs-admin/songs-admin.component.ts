import { Component, OnInit, Input } from '@angular/core';
import { Song, SongsService } from '../../../core';

@Component({
  selector: 'app-songs-admin',
  templateUrl: './songs-admin.component.html',
  styleUrls: ['./songs-admin.component.css']
})
export class SongsAdminComponent implements OnInit {
  constructor(
    private songsService: SongsService) { }

  results: Song[];

  ngOnInit() {
    this.results = [];
    this.songsService.getAllSongsAdmin().subscribe(data => {
      this.results = data;
    })
  }
}
