func ArticleFavorite(c *gin.Context) {
	slug := c.Param("slug")
	articleModel, err := FindOneArticle(&ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("Invalid slug")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	err = articleModel.favoriteBy(GetArticleUserModel(myUserModel))
	serializer := ArticleSerializer{c, articleModel}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}

func ArticleUnfavorite(c *gin.Context) {
	slug := c.Param("slug")
	articleModel, err := FindOneArticle(&ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("Invalid slug")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	err = articleModel.unFavoriteBy(GetArticleUserModel(myUserModel))
	serializer := ArticleSerializer{c, articleModel}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}