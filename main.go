package main

import (
	auth "com.lh.auth/src"
	"com.lh.service/src/tools"
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	tools.Platform(dir)
	db, err := badger.Open(badger.DefaultOptions("d:/configuration/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		auth.Test()
		c.String(http.StatusOK, "hello world!")
	})
	router.Run("localhost:8000")
}
