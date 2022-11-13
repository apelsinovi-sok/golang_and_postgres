package main

import (
	"Server/DB"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

type User struct {
	ID        uuid.UUID
	Firstname string `json:"firstname"`
	Age       uint   `json:"age"`
	Created   time.Time
}

func getInfo(ctx *gin.Context) {
	user := &User{}
	db.First(user)
	_, err := fmt.Fprintf(ctx.Writer, user.Firstname)
	if err != nil {

	}
}

func setInfo(ctx *gin.Context) {

	user := User{
		ID:        uuid.New(),
		Firstname: "Boris",
		Age:       23,
		Created:   time.Now(),
	}
	db.Create(&user)
}

var db = DB.New()

func main() {

	server := gin.Default()
	server.GET("/info", getInfo)
	server.GET("/set", setInfo)

	err := server.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
