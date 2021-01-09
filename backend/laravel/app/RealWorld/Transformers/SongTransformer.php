<?php

namespace App\RealWorld\Transformers;

class SongTransformer extends Transformer
{
    protected $resourceName = 'song';

    public function transform($data)
    {
        return [
            'id'              => $data['id'],
            'title'             => $data['title'],
            'artist'        => $data['artist'],
            'release_date'              => $data['release_date'],
            'album'           => $data['album'],
            'createdAt'         => $data['created_at']->toAtomString(),
            'updatedAt'         => $data['updated_at']->toAtomString(),
            'duration'         => $data['duration'],
            'genre'    => $data['genre']
        ];
    }
}