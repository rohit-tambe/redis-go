package main

import (
	"log"

	"github.com/go-crash-course/redis-go/pubsub"
)

var svc = pubsub.New(&pubsub.NewInput{
	RedisURL: ":6379",
})

func main() {
	Subscribe()
}

// Subscribe subscribe channel
func Subscribe() {
	// channel := fmt.Sprintf("test %s", time.Now().Add(10*time.Second).String())
	reply := make(chan []byte)
	log.Println("call subscribe method")
	err := svc.Subscribe("redis", reply)
	if err != nil {
		log.Fatal(err)
	}
}
