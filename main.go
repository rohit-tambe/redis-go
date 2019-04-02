package main

import (
	"fmt"
	"github.com/go-crash-course/redis-go/pubsub"
	"log"
)

var svc = pubsub.New(&pubsub.NewInput{
	RedisURL: ":6379",
})

func main() {
	// Publish()
	Subscribe()
}

// Publish publish channel with message
func Publish() {

	err := svc.Publish("redis", "rohit")
	if err != nil {
		log.Fatal(err)
	}
}

// Subscribe subscribe channel
func Subscribe() {
	// channel := fmt.Sprintf("test %s", time.Now().Add(10*time.Second).String())
	val := "rohit"
	reply := make(chan []byte)
	log.Println("call subscribe method")
	err := svc.Subscribe("redis", reply)
	if err != nil {
		log.Fatal(err)
	}
	msg := <-reply
	log.Println("Message from channel ", msg)
	if string(msg) != val {
		log.Fatal("expected correct reply message")
	}
	fmt.Println(msg)

}
