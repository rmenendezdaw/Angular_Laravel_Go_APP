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


    return this.apiService.get('/songs', 'go_be');
  }// end_query

  get(slug): Observable<Song> {
    return this.apiService.get('/song/' + slug)
      .pipe(map(data => data.song));
  }// end_get
  destroy(slug) {
    return this.apiService.delete('/song/' + slug);
  }
}
