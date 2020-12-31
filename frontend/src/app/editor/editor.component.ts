import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import { Song, SongsService } from '../core'

@Component({
  selector: 'app-editor-page',
  templateUrl: './editor.component.html'
})
export class EditorComponent implements OnInit {
  song: Song = {} as Song;
  songForm: FormGroup;
  errors: Object = {};
  isSubmitting = false;

  constructor(
    private songsService: SongsService,
    private route: ActivatedRoute,
    private fb: FormBuilder,
    private router: Router,
  ) {
    // use the FormBuilder to create a form group
    this.songForm = this.fb.group({
      title: '',
      artist: '',
      release_date: '',
      album: '',
      duration: '',
      genre: '',
    });

    // Optional: subscribe to value changes on the form
    // this.articleForm.valueChanges.subscribe(value => this.updateArticle(value));
  }

  ngOnInit() {
    // If there's an article prefetched, load it
    this.route.data.subscribe((data: { song: Song }) => {
      if (data.song) {
        this.song = data.song;
        this.songForm.patchValue(data.song);
      }
    });
  }

  submitForm() {
    this.isSubmitting = true;

    // update the model
    this.updateSong(this.songForm.value);


    // post the changes
    this.songsService.save(this.song).subscribe(
      data => {
        console.log(data);
        this.router.navigateByUrl('/admin')
      }, 
      err => {
        this.errors = err;
        this.isSubmitting = false;
      }
    );
  }

  updateSong(values: Object) {
    Object.assign(this.song, values);
  }
}
