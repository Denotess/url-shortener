package handlers

import (
	"database/sql"
	"url-shortener/helpers"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func Shorten(ctx *gin.Context) {
	var body models.Body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "bad json"})
		return
	}

	result, err := db.Exec(`INSERT INTO urls(short, original) VALUES (?, ?)`, "", body.Original)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "db insert failed"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to get insert id"})
		return
	}

	short, err := helpers.GenerateShortUrl(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate short code"})
		return
	}

	if _, err := db.Exec(`UPDATE urls SET short = ? WHERE id = ?`, short, id); err != nil {
		ctx.JSON(500, gin.H{"error": "db update failed"})
		return
	}

	ctx.JSON(200, gin.H{"short": short})
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
