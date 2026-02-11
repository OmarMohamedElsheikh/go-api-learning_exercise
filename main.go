package main 

import (
        "net/http"

        "github.com/gin-gonic/gin"

        "strconv"

		"database/sql"

	_	"github.com/go-sql-driver/mysql"

		"log"
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
		a.IndentedJSON(400,gin.H{"message": `you should follow this structure {
			title : album_title,
			artist : album_artist,
			price : non-negative price 
		}`})
		return
	}
	if newAlbum.Price <= 0 || newAlbum.Price > 1000 {
		a.IndentedJSON(400, gin.H{"message": "please enter a valid price"})
		return
	}

	newAlbum.ID = 
	
	
	
	a.IndentedJSON(http.StatusCreated,newAlbum)
}


func getalbID(a *gin.Context) {
	id := a.Param("id")
	for _, al := range albums {
		if al.ID == id {
			a.IndentedJSON(http.StatusOK,al)
			return
		}
	}
	a.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found!"})
}
