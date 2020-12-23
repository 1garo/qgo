package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
// TODO: try to implement something similiar to this

// implementation of remove function from container/list lib to avoid memory leak
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}
// pub function that is called
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}
*/

type qMessage struct {
	USER  string `json:"USER" binding:"required"`
	EMAIL string `json:"EMAIL" binding:"required"`
	UUID  string `json:"UUID" binding:"required"`
	MSG   string `json:"MSG" binding:"required"`
}

var queue []qMessage

func enqueue(queue []qMessage, element qMessage) []qMessage {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []qMessage) []qMessage {
	// implement the todo to remove memory leak below
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}

func CreateQueue(c *gin.Context) {
	queue = make([]qMessage, 0)
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
