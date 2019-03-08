package main

import (
	"github.com/gin-gonic/gin"
	"watchmen/config"
	"watchmen/services"
)

type Book struct {
	baba   string
	ganush int
}

func changer(arr []int) []int {
	baba := arr[1]
	arr[1] = arr[0]
	arr[0] = baba
	return arr
}
func bookChanger(book Book) {
	book.baba = "Yo"

}
func main() {
	config.Init(services.GetEnv())
	//db.Init()
	//server.Init()
}
func babmain() {
	baba := Book{baba: "aa", ganush: 12}
	myMap := map[Book]string{baba: "yo"}

	services.Logger.Info(myMap)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	services.Logger.Info(1, baba)
	bookChanger(baba)
	services.Logger.Info(2, baba)
	balance := []int{1, 2, 3}
	changer(balance)

	services.Logger.Info(1, balance)
	//services.Logger.Info(2,balance2)

	err := r.Run()
	if err != nil {
		services.Logger.Error(err)
	}
}
