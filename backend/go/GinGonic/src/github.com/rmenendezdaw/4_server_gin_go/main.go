package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// for database simulation
var lastUserId int64
var usersInDatabase = map[int64]User{}
// ..

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "HomePage"})
	fmt.Println("Endpoint homePage")
}

func handleView(c *gin.Context) {
	id := c.Query("id") // return value for id query param
	if id == "" {
		log.Println("Url Param 'id' is missing")
		return
	}
	// simulate search by id and return the user for this id
	idInt, _ := strconv.ParseInt(id, 10, 64)
	user := usersInDatabase[idInt]
	//
	c.JSON(200, user)
}

func handleViewAll(c *gin.Context) {
	// simulate search all and return the list of users
	var users []User
	for _, value := range usersInDatabase {
		users = append(users, value)
	}
	//
	c.JSON(200, users)
}

func handleSave(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	// simulate save user
	lastUserId++
	user.Id = lastUserId
	usersInDatabase[lastUserId] = user
	//
	// return the saved user
	c.JSON(200, user)
}

func main() {
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/view", handleView)
	r.GET("/all", handleViewAll)
	r.POST("/save", handleSave)

	r.Run()  // default server port 8080
}