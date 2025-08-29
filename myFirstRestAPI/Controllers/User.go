package Controllers

import (
	"firstAPI/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this one is GetUser which gets all the users present in the DataBase
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

// this is to add a new user in the database the actual logic will be performed by the model
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// to get user by its id
func GetUserByID(c *gin.Context) {
	// there will be id in the parameters of the url becuase I am getting the user using the id so the id will be present in the url itself
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound) //
	} else {
		c.JSON(http.StatusOK, user) //
	}
}

// to update the information of the user
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id) // the data will be added to the user struct
	if err != nil {
		c.JSON(http.StatusNotFound, user) // ?? here we can't abort so we saved it into the json
	}
	c.BindJSON(&user) // this will bind the new data with the user struct and the older data will be modified so in a way the data is updated
	err1 := Models.UpdateUser(&user, id)
	if err1 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// to delete any user present in the database
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"}) // gin.H is the shortcut from creating json objects
	}
}
