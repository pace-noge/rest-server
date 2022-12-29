package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pace-noge/rest-server/internal/routes"
)

func main() {
	router := gin.Default()
	server := routes.NewTaskServer()

	router.POST("/task/", server.CreateTaskhandler)
	router.GET("/task/", server.GetAllTaskHandler)
	router.DELETE("/task/", server.DeleteAllTaskHandler)
	router.GET("/task/:id", server.GetTaskHandler)
	router.DELETE("/task/:id", server.DeleteTaskHandler)
	router.GET("/tag/:tag", server.TagHandler)
	router.GET("/due/:year/:month/:day", server.DueHandler)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))

}
