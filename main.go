package main

import (
	"fmt"
	h "github.com/kooksee/gin-demo/h"
	app "github.com/kooksee/gin-demo/app"
	utils "github.com/kooksee/gin-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"net/http"

	"log"
)

func main() {

	h.InitHello()

	s := app.GetInstance()
	if h1, ok := s.Get("h").(*h.Hello); ok {
		h1.Say("hello")
		fmt.Println(h1.GetName())
	}

	utils.InitRedis()
	utils.StringOperation()
	utils.ListOperation()
	utils.SetOperation()
	utils.HashOperation()
	utils.ConnectPool()



	//ssssssssssss
	r := gin.Default()
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			m := make(map[string]interface{})
			m["a"] = "你好"
			e := so.Emit("cn1111", m)
			//这个没有问题
			fmt.Println("\n\n")

			b := make(map[string]string)
			b["u-a"] = "中文内容" //这个不能是中文
			m["b-c"] = b
			e = so.Emit("cn2222", m)
			log.Println(e)

			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		// Socket.io acknowledgement example
		// The return type may vary depending on whether you will return
		// For this example it is "string" type
		so.On("chat message with ack", func(msg string) string {
			return msg
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	r.GET("/1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "asset")

	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})



	//r.Handle("/socket.io/", server)
	//r.Handle("/", http.FileServer(http.Dir("./asset")))
	//r.Run() // listen and serve on 0.0.0.0:8080

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	http.Handle("/socket.io/", server)

	srv.ListenAndServe()


	//log.Println("Serving at localhost:5000...")
	//log.Fatal(http.ListenAndServe(":5000", nil))
}



