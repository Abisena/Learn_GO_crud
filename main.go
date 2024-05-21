package main

import (
	"fmt"
	"go-app/config"
	"go-app/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.GET("/getUser", gin.WrapH(controller.GetUserHandler(db)))
	r.GET("/getUser/:id", gin.WrapH(controller.GetUserByID(db)))
	r.POST("/createUser", gin.WrapH(controller.CreateUserHandler(db)))
	r.PUT("/updateUser/:id", gin.WrapH(controller.UpdateUserHandler(db)))
	r.DELETE("/deleteUser/:id", gin.WrapH(controller.DeleteUserHandler(db)))

	r.Run("localhost:8000")
	fmt.Println("Hello GoNers")
}