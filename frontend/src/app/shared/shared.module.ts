import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';

import { ArticleListComponent, ArticleMetaComponent, ArticlePreviewComponent } from './article-helpers';
import { SongsComponent, SongsPreviewComponent } from './songs-helpers';
import { FavoriteButtonComponent, FollowButtonComponent } from './buttons';
import { ListErrorsComponent } from './list-errors.component';
import { ShowAuthedDirective } from './show-authed.directive';
import { SongsAdminComponent } from './admin-helpers/songs-admin/songs-admin.component';
import { UsersAdminComponent } from './admin-helpers/users-admin/users-admin.component';
import { SongsAdminPreviewComponent } from './admin-helpers/songs-admin';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule
  ],
  declarations: [
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    FavoriteButtonComponent,
    FollowButtonComponent,
    ListErrorsComponent,
    ShowAuthedDirective,
    SongsComponent,
    SongsPreviewComponent,
    SongsAdminComponent,
    SongsAdminPreviewComponent,
    UsersAdminComponent

  ],
  exports: [
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    CommonModule,
    FavoriteButtonComponent,
    FollowButtonComponent,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    ListErrorsComponent,
    RouterModule,
    ShowAuthedDirective,
    SongsComponent,
    SongsPreviewComponent,
    SongsAdminComponent,
    SongsAdminPreviewComponent,
    UsersAdminComponent
  ]
})
export class SharedModule {}
