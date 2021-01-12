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
		r.GET("/dqueue", ctl.RemoveElemFromQueue)
		r.GET("/dqueueLast", ctl.RemoveLastElemFromQueue)
		r.GET("/remove", ctl.RemoveElemFromQueue)
		r.POST("/query", ctl.GetQueues)
	}
	err := r.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
