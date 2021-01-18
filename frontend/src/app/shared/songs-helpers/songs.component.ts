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
    @Input() limit: number;
    @Input()
    set config(config: SongListConfig) {
      if (config) {
        this.query = config;
        this.currentPage = 1;
        this.runQuery();
      }
    }

  query: SongListConfig
  results: Song[];
  loading = false;
  currentPage = 1;
  songsCount = 0;
  totalPages: Array<number> = [1];

  setPageTo(pageNumber) {
    this.currentPage = pageNumber;
    this.runQuery();
  }
  runQuery() {
    this.loading = true;
    this.results = [];
 
    if (this.views) {
      this.query.filters.views = this.views;
    }

    if (this.release_date) {
      this.query.filters.release_date = this.release_date;
    }
    if (this.limit) {
      this.query.filters.limit = this.limit;
      this.query.filters.offset =  (this.limit * (this.currentPage - 1));
    }

    console.log(this.query);
    this.songsService.query(this.query)
    .subscribe(data => {
      console.log(data);
      this.loading = false;
      
      if (this.query.type == "favorites") {
          data.songs.forEach(t => {
            // console.log(t);
            if (t.favorited == true) this.results.push(t)
          })
      }else this.results = data.songs;

      this.totalPages = Array.from(new Array(Math.ceil(data.songsCount / this.limit)), (val, index) => index + 1);

    });
  }

}
