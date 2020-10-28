package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Println("Request in") // ①
		c.Next() // next handler func
		log.Println("Response out") // ②
	})

	r.Use(func(c *gin.Context) {
		log.Println("Request2 in") // ①
		c.Next() // next handler func
		log.Println("Response2 out") // ②
	})

	r.GET("/ping", func(c *gin.Context) {
		log.Println("ping")
		c.String(http.StatusOK, "%s", "pong!")
	})

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalln(err)
	}
}

// 2019/10/08 17:18:06 Request in
// 2019/10/08 17:18:06 Request2 in
// 2019/10/08 17:18:06 ping
// 2019/10/08 17:18:06 Response2 out
// 2019/10/08 17:18:06 Response out
// [GIN] 2019/10/08 - 17:18:06 | 200 |     318.906µs |       127.0.0.1 | GET      /ping
