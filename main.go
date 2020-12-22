package main

import (
	"container/list"
	"log"
	ctl "qgo/controller"

	"github.com/gin-gonic/gin"
)

func createQueue() *list.List {
	return list.New()
}

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
