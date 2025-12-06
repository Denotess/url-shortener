package handlers

import (
	"database/sql"
	"url-shortener/helpers"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func Shorten(ctx *gin.Context) {
	var body models.Body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "bad json"})
		return
	}

	short, err := helpers.GenerateShortUrl(body.Original)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})
	}

	if _, err := db.Exec(`INSER INTO urls(short, original) VALUES (?, ?)`, short, body.Original); err != nil {
		ctx.JSON(500, gin.H{"error": "db insert failed"})
		return
	}
}

func Redirect(ctx *gin.Context) {
	short := ctx.Param("short")
	var original string
	if err := db.QueryRow(`SELECT original FROM urls WHERE short = ?`, short).Scan(&original); err != nil {
		ctx.JSON(404, gin.H{"error": "not found"})
		return
	}
	ctx.Redirect(302, original)
}
