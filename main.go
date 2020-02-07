package main

import (
	"GinDemo/Api"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Adduser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "you are log in ")

}
func AdduserForm(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "you are log in ")
}
func initDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	return db
}
func main() {
	fmt.Println("welcome to go gin ")
	db := initDB()
	defer db.Close()
	Api.CreateApiEndpoints()
	err := Api.R.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
