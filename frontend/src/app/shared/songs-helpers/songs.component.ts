import { Component, OnInit, Input } from '@angular/core';
import { Song, SongsService } from '../../core';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.css']
})
export class SongsComponent implements OnInit {
  constructor(
    private songsService: SongsService) { }

  ngOnInit() {
    console.log('holi');
    this.songsService.query().subscribe(data => {
      console.log(data);
    })
  }

}
