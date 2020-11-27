import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SongComponent } from './song.component';
import { SongResolver } from './song-resolver.service';

const routes: Routes = [
  {
    path: ':id',
    component: SongComponent,
    resolve: {
      song: SongResolver
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SongRoutingModule {}
