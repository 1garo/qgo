package controller

import (
	"container/list"
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

var queue *list.List

func CreateQueue(c *gin.Context) {
	queue = list.New()
	c.JSON(http.StatusOK, gin.H{
		"msg": "queue was created successfully!!!",
	})
}

func PopulateQueues(c *gin.Context) {
	if queue == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue doesn't exist, please create it!!!",
		})
	} else {
		queue.PushFront(qMessage{
			USER:  "roberto",
			EMAIL: "roberto@rr.com",
			UUID:  "InternalID:1",
			MSG:   "lindo",
		})
		queue.PushFront(qMessage{
			USER:  "alex",
			EMAIL: "alex@rr.com",
			UUID:  "InternalID:2",
			MSG:   "lindox",
		})
		queue.PushBack(qMessage{
			USER:  "ale",
			EMAIL: "ale@rr.com",
			UUID:  "InternalID:3",
			MSG:   "linduxo",
		})
		c.JSON(http.StatusOK, gin.H{
			"msg": queue.Front().Value,
		})
	}
}

func GetQueues(c *gin.Context) {
	size := queue.Len()
	log.Printf("squeue: %v", queue)
	resp := make([]interface{}, size)
	if size == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue don't have any item!",
		})
		return
	}
	valueBack := queue.Back()
	if valueBack != nil {
		resp = append(resp, valueBack.Value)
		// queue.Remove(valueBack)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue don't have any back item!",
		})
	}
	valueFront := queue.Front().Next().Value
	if valueFront != nil {
		resp = append(resp, valueFront)
		// queue.Remove(valueFront)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "queue don't have any front item!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"queues": resp,
	})
	queue.Remove(valueBack)
	queue.Remove(queue.Front())
	log.Printf("equeue: %v", queue)
}
