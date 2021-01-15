import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { Song, SongListConfig } from '../models';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class SongsService {

  constructor(private apiService: ApiService) { 
  }

  query(query: SongListConfig): Observable<{songs: Song[], songsCount: number}> {
    const params = {};

    Object.keys(query.filters)
    .forEach((key) => {
      params[key] = query.filters[key];
    });

    console.log(params);

    return this.apiService.get('/songs', 'api_url', new HttpParams({ fromObject: params }))
    .pipe(map(data => data.songs));
  }// end_query

  get(id): Observable<Song> {
    return this.apiService.get('/songs/' + id)
      .pipe(map(data => data.song));
  }// end_get
  destroy(id) {
    return this.apiService.delete('/songs/' + id);
  }

  getLaravel(id): Observable<Song> {
    return this.apiService.get('/songs/' + id, 'laravel_be')
      .pipe(map(data => data.song));
  }// end_get

  getAllSongsAdmin(): Observable<Song[]> {
    return this.apiService.get('/songs', 'laravel_be').pipe(map(data => data.songs));;
  }// end_getAllSongsAdmin

  save(song): Observable<Song> {
    if (song.id) {
      return this.apiService.putLaravel('/songs/' + song.id, {song: song})
        .pipe(map(data => data));

    } else {
      return this.apiService.postLaravel('/songs/', {song: song})
        .pipe(map(data => data));
    }
  }

  favorite(id): Observable<Song> {
    return this.apiService.post('/songs/' + id + '/favorite');
  }

  unfavorite(id): Observable<Song> {
    return this.apiService.deleteGo('/songs/' + id + '/favorite');
  }
}
