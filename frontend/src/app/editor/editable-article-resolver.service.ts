import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { Song, SongsService, UserService } from '../core';
import { catchError ,  map } from 'rxjs/operators';

@Injectable()
export class EditableArticleResolver implements Resolve<Song> {
  constructor(
    private songsService: SongsService,
    private router: Router,
    private userService: UserService
  ) { }

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {

    return this.songsService.getLaravel(route.params['slug'])
      .pipe(
        map(
          song => {
              return song;
          }
        ),
        catchError((err) => this.router.navigateByUrl('/'))
      );
  }
}
