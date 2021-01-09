<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Http\Requests\Api\CreateSong;

use Redis;
use App\Song;

class SongController extends Controller {
    public function store(CreateSong $request) {
        
        $song = new Song();
        
        $song -> title = $request['song']['title'];
        $song -> artist = $request['song']['artist'];
        $song -> release_date = $request['song']['release_date'];
        $song -> album = $request['song']['album'];
        $song -> duration = $request['song']['duration'];
        $song -> genre = $request['song']['genre'];
        
        $song -> save();

        return response() -> json($song);

    }// end_create

    public function index() {
        $songs = Song::all();

        return response() -> json($songs);
    }// end_show

    public function show($id) {
        $song = Song::find($id);

        $response = ["song" => $song];

        return response() -> json($response);
    }// end_showSong

    public function update(Request $request, $id) {
        $song = Song::find($id);

        if (!$song) return response() -> json('Not Found');

        $song -> title = $request['song']['title'];
        $song -> artist = $request['song']['artist'];
        $song -> release_date = $request['song']['release_date'];
        $song -> album = $request['song']['album'];
        $song -> duration = $request['song']['duration'];
        $song -> genre = $request['song']['genre'];

        $song -> save();

        return response() -> json($song);
    }// end_update

    public function destroy($id) {
        $song = Song::find($id);
        
        if (!$song) return response() -> json('Not Found');

        $song -> delete();

        return response() -> json($song);
    }// end_delete
}// end_SongController