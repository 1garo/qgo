package main

import (
	"log"
	ctl "qgo/controller"

	"github.com/gin-gonic/gin"
)

// TODO: implement the concept of pub/sub
// pub being who post the msg
// sub being the consumer

func main() {

	// Dequeue
	//front := queue.Front()
	//fmt.Println(front.Value)
	//fmt.Printf("%v", queue.Back().Value)
	//// This will remove the allocated memory and avoid memory leaks
	//queue.Remove(front)
	r := gin.Default()
	_ = r.Group("queue")
	{
		r.GET("/create", ctl.CreateQueue)
		r.GET("/populate", ctl.PopulateQueues)
		r.POST("/query", ctl.GetQueues)
	}
	err := r.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
