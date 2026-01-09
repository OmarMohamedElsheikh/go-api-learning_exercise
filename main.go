package main 

import (
        "net/http"

        "github.com/gin-gonic/gin"

        "strconv"

)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}


var albums = []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()
	r.GET("/albums", get_alb)
	r.GET("/albums/:id",getalbID)
	r.POST("/albums", post_alb)
	r.Run("localhost:8080")
}


func get_alb(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, albums)
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

	newAlbum.ID = strconv.Itoa(len(albums) + 1)
	
	albums = append(albums, newAlbum)
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
