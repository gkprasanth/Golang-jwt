package main

import (
	"gin-backend/controllers"
	"gin-backend/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
	//initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to gin application",
		})
	})

	router.POST("/signup", controllers.Signup2)
	router.POST("/login", controllers.Login1)
	router.Run(":4444")
}
