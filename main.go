package main

import (
	"database/sql"
	"log"
	"net/http"
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short TEXT UNIQUE NOT NULL,
		original TEXT NOT NULL
	);`); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handlers.SetDB(db)

	router.POST("/shorten", handlers.Shorten)
	router.GET("/:short", handlers.Redirect)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
