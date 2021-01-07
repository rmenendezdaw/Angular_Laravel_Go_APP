import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import {
  Song,
  SongsService,
  CommentsService,
  User,
  UserService
} from '../core';

@Component({
  selector: 'app-song-page',
  templateUrl: './song.component.html'
})
export class SongComponent implements OnInit {
  song: Song;
  currentUser: User;
  canModify: boolean;
  isSubmitting = false;
  isDeleting = false;

  constructor(
    private route: ActivatedRoute,
    private songsService: SongsService,
    private commentsService: CommentsService,
    private router: Router,
    private userService: UserService,
  ) { }

  ngOnInit() {
    // Retreive the prefetched song
    this.route.data.subscribe(
      (data: { song: Song }) => {
        this.song = data.song;
        // Load the comments on this song
      }
    );

    // Load the current user's data
    this.userService.currentUser.subscribe(
      (userData: User) => {
        this.currentUser = userData;

      }
    );
  }
}
