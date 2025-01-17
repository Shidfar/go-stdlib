package legacy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	type object struct {
		Key          string
		LastModified time.Time
	}

	rand.Seed(time.Now().UnixNano())
	listSize := rand.Intn(2001) + 1000
	rand.Seed(time.Now().UnixNano())
	priorityMaxSize := rand.Intn(256) + 127
	now := time.Now()
	var objectList []object
	for i := 1; i < listSize; i++ {
		objectList = append(objectList, object{
			Key:          fmt.Sprintf("object %d", i),
			LastModified: now.Add(time.Duration(i) * time.Second),
		})
	}

	priorityQueue := NewPriorityQueue(priorityMaxSize)
	for i, obj := range objectList {
		priorityQueue.Push(Item{
			Key:          obj.Key,
			LastModified: obj.LastModified,
			Index:        i,
		})
	}

	for priorityQueue.Len() > 0 {
		item := priorityQueue.Pop().(Item)
		for i := 0; i < len(objectList)-priorityMaxSize; i++ {
			if objectList[i].Key == item.Key {
				fmt.Println("max pq capacity: ", priorityMaxSize, " obj list length ", listSize)
				fmt.Println(item.Key, " should not be in part of priority queue")
				t.Fatal("priority queue is broken")
			}
		}
	}
}
