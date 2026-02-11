package main 

import (
        "net/http"

        "github.com/gin-gonic/gin"

		"database/sql"

	_	"github.com/go-sql-driver/mysql"

		"log"

		"fmt"
)



type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}


func main() {

	db, err := sql.Open("mysql", 
	           "root@tcp(127.0.0.1:3306)/recordings")
	 if err != nil {
	     log.Fatal("Error connecting to the database:", err)
	 }
	 defer db.Close()

	 err = db.Ping()
	 if err != nil {
	     log.Fatal("Error pinging the database:", err)
	 }

	 log.Println("Successfully connected to the database!")
}
	r := gin.Default()
	r.GET("/albums", get_alb)
	r.GET("/albums/:id",getalbID)
	r.POST("/albums", post_alb)
	r.Run("localhost:8080")
}

func get_alb(c *gin.Context){
	
	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var albums []Album

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, alb)
	}

	c.IndentedJSON(http.StatusOK, albums)
	}

func post_alb(a *gin.Context) {
	var newAlbum album 
	err := a.ShouldBindJSON(&newAlbum)

	if err != nil {
		a.IndentedJSON(http.StatusBadRequest,gin.H{"message": `you should follow this structure {
			title : album_title,
			artist : album_artist,
			price : non-negative price 
		}`})
		return
	}
	if newAlbum.Price <= 0 || newAlbum.Price > 1000 {
		a.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please enter a valid price"})
		return
	}

	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?,?,?)",newAlbum.Title, newAlbum.Artist,newAlbum.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("POST album: %v", err),
		})
		return
	}
	id, _ := result.LastInsertId()
	newAlbum.ID = int(id)
	
	a.IndentedJSON(http.StatusCreated,newAlbum)
}


func getalbID(a *gin.Context) {
	id := a.Param("id")

	var alb album
	
	rows, err := db.Query("SELECT id, title, artist, price FROM albums WHERE id = ?",id).Scan(&alb.ID,&alb.Title,&alb.Artist,&alb.Price)

	if err != nil {
		if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "album not found",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
	defer rows.Close()
	a.JSON(http.StatusOK,rows)
}
