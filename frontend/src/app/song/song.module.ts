import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { SongComponent } from './song.component';
import { SongResolver } from './song-resolver.service';
import { SharedModule } from '../shared';
import { SongRoutingModule } from './song-routing.module';

@NgModule({
  imports: [
    SharedModule,
    SongRoutingModule
  ],
  declarations: [
    SongComponent,    
  ],

  providers: [
    SongResolver
  ]
})
export class SongModule {}
