package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "go-mapnik"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func makemap() {
    m := mapnik.New()
    if err := m.Load("./src/data/map.xml"); err != nil {
	log.Fatal(err)
    }
    m.Resize(1000, 500)
    m.ZoomTo(-180, -90, 180, 90)
    opts := mapnik.RenderOpts{Format: "png32"}
    if err := m.RenderToFile(opts, "./src/example-1.png"); err != nil {
	log.Fatal(err)
    }
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    makemap()
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.Run("0.0.0.0:8080")
}
