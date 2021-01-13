package main

import (
	"log"
	ctl "qgo/controller"

	"github.com/gin-gonic/gin"
)

/*
TODO: implement the concept of pub/sub
- pub being who post the msg
- sub being the consumer
*/

func main() {
	r := gin.Default()
	_ = r.Group("queue")
	{
		r.GET("/create", ctl.CreateQueue)
		r.GET("/populate", ctl.PopulateQueues)
		r.GET("/dqueue", ctl.RemoveElemFromBottom)
		r.GET("/dqueueLast", ctl.RemoveLastElemFromTop)
		r.GET("/remove", ctl.RemoveElemFromBottom)
		r.POST("/query", ctl.GetQueues)
	}
	err := r.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
