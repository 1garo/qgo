package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type qMessage struct {
	USER  string `json:"USER" binding:"required"`
	EMAIL string `json:"EMAIL" binding:"required"`
	UUID  string `json:"UUID" binding:"required"`
	MSG   string `json:"MSG" binding:"required"`
}

var queue []qMessage = nil

func enqueue(queue []qMessage, element qMessage) []qMessage {
	queue = append(queue, element)
	log.Println("Enqueued: ", element)
	return queue
}

func dequeueTop(queue []qMessage) []qMessage {
	size := len(queue) - 1
	element := queue[size]
	// this garbage collector collect this
	// and not leak memory
	queue[size] = qMessage{}
	log.Println("Dequeued from top: ", element)
	return queue[:size-1]
}

func dequeueBottom(queue []qMessage) []qMessage {
	element := queue[0]
	queue[0] = qMessage{}
	log.Println("Dequeued from bottom: ", element)
	return queue[1:]
}

// CreateQueue -> create empty queue object
func CreateQueue(c *gin.Context) {
	queue = make([]qMessage, 0)
	c.JSON(http.StatusOK, gin.H{
		"msg": "queue was created successfully!!!",
	})
}

// RemoveElemFromBottom -> remove element from the bottom of the queue
func RemoveElemFromBottom(c *gin.Context) {
	if queue == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue doesn't exist, please create it!!!",
		})
		return
	}
	queue = dequeueBottom(queue)
	c.JSON(http.StatusOK, gin.H{
		"msg": queue,
	})
}

// RemoveLastElemFromTop -> remove element from the top of the queue
func RemoveLastElemFromTop(c *gin.Context) { dequeueTop([]qMessage{}) }

// PopulateQueues -> dummy function to create data for test purpouse
func PopulateQueues(c *gin.Context) {
	if queue == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue doesn't exist, please create it!!!",
		})
		return
	}
	queue = enqueue(queue, qMessage{
		USER:  "roberto",
		EMAIL: "roberto@rr.com",
		UUID:  "1",
		MSG:   "lindo",
	})
	queue = enqueue(queue, qMessage{
		USER:  "alex",
		EMAIL: "alex@rr.com",
		UUID:  "2",
		MSG:   "lindox",
	})
	queue = enqueue(queue, qMessage{
		USER:  "ale",
		EMAIL: "ale@rr.com",
		UUID:  "3",
		MSG:   "linduxo",
	})
	c.JSON(http.StatusOK, gin.H{
		"msg": queue,
	})
}

// GetQueues -> return all queues
func GetQueues(c *gin.Context) {
	//TODO: create a while both back and front until value is != nil
	size := len(queue)
	log.Printf("squeue: %v", queue)
	if size == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue don't have any item!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"queues": queue,
	})
	log.Printf("equeue: %v", queue)
}
