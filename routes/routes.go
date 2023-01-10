package routes

import (
	"github.com/gin-gonic/gin"
	"repos/project/controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//grp1 := r.Group("/employee-api")
	r.GET("/user", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	//grp1.GET("/employee", controllers.GetUsers)
	//grp1.GET("employee/:id", controllers.GetUserByID)
	//grp1.PUT("employee/:id", controllers.UpdateUser)
	//grp1.DELETE("employee/:id", controllers.DeleteUser)

	return r
}
