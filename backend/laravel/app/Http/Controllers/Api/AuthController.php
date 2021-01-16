<?php

namespace App\Http\Controllers\Api;

use Auth;
use App\User;
use App\Http\Requests\Api\LoginUser;
use App\Http\Requests\Api\RegisterUser;
use App\RealWorld\Transformers\UserTransformer;

use Redis;

class AuthController extends ApiController
{
    /**
     * AuthController constructor.
     *
     * @param UserTransformer $transformer
     */
    public function __construct(UserTransformer $transformer)
    {
        $this->transformer = $transformer;
    }

    /**
     * Login user and return the user if successful.
     *
     * @param LoginUser $request
     * @return \Illuminate\Http\JsonResponse
     */

    public function getLogins() {
        return response() -> json(Redis::get('currentUsers'));
    }// end_getLogins


    public function login(LoginUser $request)
    {
        $credentials = $request->only('user.email', 'user.password');
        $email = $request["user"]["email"];
        $credentials = $credentials['user'];

        // echo Redis::get("user");
        // echo Redis::get("token");

        if ($email != Redis::get("user")) {
            return $this -> respondFailedLogin();
        }

        if (! Auth::once($credentials)) {
            return $this->respondFailedLogin();
        }

        return $this->respondWithTransformer(auth()->user());
    }

    /**
     * Register a new user and return the user if successful.
     *
     * @param RegisterUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function register(RegisterUser $request)
    {
        $user = User::create([
            'username' => $request->input('user.username'),
            'email' => $request->input('user.email'),
            'password' => bcrypt($request->input('user.password')),
        ]);

        return $this->respondWithTransformer($user);
    }
}
