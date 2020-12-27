import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { Song } from '../models';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class SongsService {

  constructor(private apiService: ApiService) { 
  }

  query(): Observable<Song[]> {
    const params = {};


    return this.apiService.get('/songs');
  }// end_query

  get(id): Observable<Song> {
    return this.apiService.get('/songs/' + id)
      .pipe(map(data => data.song));
  }// end_get
  destroy(id) {
    return this.apiService.delete('/songs/' + id);
  }

  getAllSongsAdmin(): Observable<Song[]> {
    return this.apiService.get('/songs', 'laravel_be');
  }// end_getAllSongsAdmin

  save(song): Observable<Song> {
    if (song.id) {
      return this.apiService.putLaravel('/songs/' + song.id, {song: song})
        .pipe(map(data => data.song));

    } else {
      return this.apiService.postLaravel('/songs/', {song: song})
        .pipe(map(data => data.song));
    }
  }
}
