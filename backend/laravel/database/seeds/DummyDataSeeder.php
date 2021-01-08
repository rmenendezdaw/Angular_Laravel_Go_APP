<?php

use Illuminate\Database\Seeder;

class DummyDataSeeder extends Seeder
{
    /**
     * Total number of users.
     *
     * @var int
     */
    protected $totalUsers = 25;

    /**
     * Total number of tags.
     *
     * @var int
     */
    protected $totalTags = 10;

    protected $totalSongs = 15;

    /**
     * Percentage of users with articles.
     *
     * @var float Value should be between 0 - 1.0
     */
    protected $userWithArticleRatio = 0.8;

    /**
     * Maximum articles that can be created by a user.
     *
     * @var int
     */
    protected $maxArticlesByUser = 15;

    /**
     * Maximum tags that can be attached to an article.
     *
     * @var int
     */
    protected $maxArticleTags = 3;

    /**
     * Maximum number of comments that can be added to an article.
     *
     * @var int
     */
    protected $maxCommentsInArticle = 10;

    /**
     * Percentage of users with favorites.
     *
     * @var float Value should be between 0 - 1.0
     */
    protected $usersWithFavoritesRatio = 0.75;

    /**
     * Percentage of users with following.
     *
     * @var float Value should be between 0 - 1.0
     */
    protected $usersWithFollowingRatio = 0.75;

    /**
     * Populate the database with dummy data for testing.
     * Complete dummy data generation including relationships.
     * Set the property values as required before running database seeder.
     *
     * @param \Faker\Generator $faker
     */
    public function run(\Faker\Generator $faker)
    {

        $songs = factory(\App\Song::class) -> times ($this -> totalSongs) -> create();
       
    }
}
