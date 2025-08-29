package Routes

import (
	"firstAPI/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter... Configure Router
func SetUpRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}
