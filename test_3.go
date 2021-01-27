package main

import (
	"github.com/gin-gonic/gin"
)
	



func main() {
	c := gin.Default()
	// c.GET("/user/:name", func(c *gin.Context){
	// 	name := c.Param("name")
	// 	c.String(200, "Hello %s" , name)
	// })

	c.GET("/user/:name/*action", func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(200, message)
	})
	c.Run()
}