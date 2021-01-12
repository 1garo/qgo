package controller

import (
	"fmt"
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
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeueTop(queue []qMessage) []qMessage { return []qMessage{} }

func dequeueBottom(queue []qMessage) []qMessage {
	element := queue[0]
	queue[0] = qMessage{}
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func CreateQueue(c *gin.Context) {
	queue = make([]qMessage, 0)
	c.JSON(http.StatusOK, gin.H{
		"msg": "queue was created successfully!!!",
	})
}

func RemoveElemFromQueue(c *gin.Context) {
	if queue == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue doesn't exist, please create it!!!",
		})
		return
	} else {
		queue = dequeueBottom(queue)
		c.JSON(http.StatusOK, gin.H{
			"msg": queue,
		})
	}
}

// TODO: implement func to remove from top instead of bottom
func RemoveLastElemFromQueue(c *gin.Context) { dequeueTop([]qMessage{}) }

func PopulateQueues(c *gin.Context) {
	if queue == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue doesn't exist, please create it!!!",
		})
		return
	} else {
		queue = enqueue(queue, qMessage{
			USER:  "roberto",
			EMAIL: "roberto@rr.com",
			UUID:  "InternalID:1",
			MSG:   "lindo",
		})
		queue = enqueue(queue, qMessage{
			USER:  "alex",
			EMAIL: "alex@rr.com",
			UUID:  "InternalID:2",
			MSG:   "lindox",
		})
		queue = enqueue(queue, qMessage{
			USER:  "ale",
			EMAIL: "ale@rr.com",
			UUID:  "InternalID:3",
			MSG:   "linduxo",
		})
		c.JSON(http.StatusOK, gin.H{
			"msg": queue,
		})
	}
}

func GetQueues(c *gin.Context) {
	//TODO: create a while both back and front until value is != nil
	size := len(queue)
	log.Printf("squeue: %v", queue)
	// resp := make([]interface{}, size)
	if size == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue don't have any item!",
		})
		return
	}
	// valueBack := queue.Back()
	// if valueBack != nil {
	// 	resp = append(resp, valueBack.Value)
	// 	// queue.Remove(valueBack)
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"msg": "queue don't have any back item!",
	// 	})
	// }
	// valueFront := queue.Front()
	// if valueFront != nil {
	// 	resp = append(resp, valueFront)
	// 	// queue.Remove(valueFront)
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"msg": "queue don't have any front item!",
	// 	})
	// }
	c.JSON(http.StatusOK, gin.H{
		"queues": queue,
	})
	// defer queue.Remove(valueBack)
	// defer queue.Remove(valueFront)
	log.Printf("equeue: %v", queue)
}
