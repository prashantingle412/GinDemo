package Api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

type Emp struct {
	Name string `json:"name,omitempty"`
}

func Hello(c *gin.Context) {
	log.Println("writer")
	c.JSON(200, gin.H{
		"messege": "hello",
	})
}
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"messege": "this is protected route",
	})
}
func Register(c *gin.Context) {
	var emp Emp
	err := c.BindJSON(&emp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, emp)
}
func CreateApiEndpoints() {
	R = gin.Default()
	v1 := R.Group("/v1")
	{
		v1.GET("/hello", Hello)
		v1.GET("/login", Login)
	}
	v2 := R.Group("/v2")
	{
		v2.POST("/register", Register)
		v2.GET("/login", Login)
	}
}
