import { Component, Input } from '@angular/core';
import { Song, SongsService, SongListConfig } from '../../core';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.css']
})
export class SongsComponent {
  constructor(
    private songsService: SongsService) { }

    @Input() views: string;
    @Input() release_date: string;
    @Input()
    set config(config: SongListConfig) {
      if (config) {
        this.query = config;
        this.runQuery();
      }
    }

  query: SongListConfig
  results: Song[];
  loading = false;

  runQuery() {
    this.loading = true;
    this.results = [];
 
    if (this.views) {
      this.query.filters.views = this.views;
    }

    if (this.release_date) {
      this.query.filters.release_date = this.release_date;
    }

    this.songsService.query(this.query)
    .subscribe(data => {
      this.loading = false;
      
      if (this.query.type == "favorites") {
          data.forEach(t => {
            if (t.favorited == true) this.results.push(t)
          })
      }else this.results = data;
    });
  }

}
