<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Song extends Model
{
    protected $table = 'songs';
    protected $filltable = ['title', 'artist', 'releaseDate', 'album', 'duration', 'genre', 'views'];
}
