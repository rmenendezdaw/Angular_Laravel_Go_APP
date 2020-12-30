<?php

namespace App\Http\Requests\Api;

class CreateSong extends ApiRequest
{
    /**
     * Get data to be validated from the request.
     *
     * @return array
     */
    protected function validationData()
    {
        return $this->get('song') ?: [];
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'title' => 'required|string|max:255',
            'artist' => 'required|string|max:255',
            'release_date' => 'required|string|max:255',
            'album' => 'required|string|max:255',
            'duration' => 'required|integer|max:255',
            'genre' => 'required|string|max:255'
        ];
    }
}
