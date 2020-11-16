<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Song;

class SongController extends Controller {
    public function create(Request $request) {
        $song = new Song();

        $song -> title = $request -> title;
        $song -> artist = $request -> artist;
        $song -> releaseDate = $request -> releaseDate;
        $song -> album = $request -> album;
        $song -> duration = $request -> duration;
        $song -> genre = $request -> genre;

        $song -> save();

        return response() -> json($song);
    }// end_create

    public function show() {
        $songs = Song::all();

        return response() -> json($songs);
    }// end_show

    public function showSong($id) {
        $song = Song::find($id);

        return response() -> json($song);
    }// end_showSong

    public function update(Request $request, $id) {
        $song = Song::find($id);

        if (!$song) return response() -> json('Not Found');

        $song -> title = $request -> title;
        $song -> artist = $request -> artist;
        $song -> releaseDate = $request -> releaseDate;
        $song -> album = $request -> album;
        $song -> duration = $request -> duration;
        $song -> genre = $request -> genre;

        $song -> save();

        return response() -> json($song);
    }// end_update

    public function delete($id) {
        $song = Song::find($id);
        
        if (!$song) return response() -> json('Not Found');

        $song -> delete();

        return response() -> json($song);
    }// end_delete
}// end_SongController