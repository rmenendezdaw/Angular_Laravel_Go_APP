<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Http\Requests\Api\CreateSong;
use App\Http\Requests\Api\UpdateSong;
use App\RealWorld\Transformers\SongTransformer;



use Redis;
use App\Song;

class SongController extends ApiController {
    public function __construct(SongTransformer $transformer)
    {
        $this -> transformer = $transformer;
    }

    public function store(CreateSong $request) {
        
        $song = new Song();
        
        $song -> title = $request['song']['title'];
        $song -> artist = $request['song']['artist'];
        $song -> release_date = $request['song']['release_date'];
        $song -> album = $request['song']['album'];
        $song -> duration = $request['song']['duration'];
        $song -> genre = $request['song']['genre'];
        
        $song -> save();

        return $this -> respondWithTransformer($song);


    }// end_create

    public function index() {
        $songs = Song::all();

        return $this -> respondWithTransformer($songs);
    }// end_show

    public function show($id) {
        $song = Song::find($id);

        $response = ["song" => $song];

        return $this -> respondWithTransformer($song);
    }// end_showSong

    public function update(UpdateSong $request, $id) {
        $song = Song::find($id);

        if (!$song) return response() -> json('Not Found');

        $song -> title = $request['song']['title'];
        $song -> artist = $request['song']['artist'];
        $song -> release_date = $request['song']['release_date'];
        $song -> album = $request['song']['album'];
        $song -> duration = $request['song']['duration'];
        $song -> genre = $request['song']['genre'];

        $song -> save();

        return $this -> respondWithTransformer($song);
    }// end_update

    public function destroy($id) {
        $song = Song::find($id);
        
        if (!$song) return response() -> json('Not Found');

        $song -> delete();

        return $this -> respondSuccess();
    }// end_delete
}// end_SongController