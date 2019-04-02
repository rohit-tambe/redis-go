package pubsub

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Service service
type Service struct {
	pool *redis.Pool
	conn redis.Conn
}

// NewInput input for constructor
type NewInput struct {
	RedisURL string
}

// New return nee service
func New(input *NewInput) *Service {
	if input == nil {
		log.Fatal("input is required")
	}
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", input.RedisURL)
		},
	}
	// Get a connection
	conn := redispool.Get()
	defer conn.Close()

	// test the connection
	_, err := conn.Do("ping")
	if err != nil {
		log.Fatalf("can't connect to the redis database, got error:\n%v", err)
	}
	return &Service{
		pool: redispool,
		conn: conn,
	}
}

// Publish publish key value
func (s *Service) Publish(key, value string) error {
	conn := s.pool.Get()
	conn.Do("publish", key, value)
	return nil
}

// Subscribe subscribe message
func (s *Service) Subscribe(key string, msg chan []byte) error {
	log.Println("In subscribe")
	rc := s.pool.Get()
	psc := redis.PubSubConn{
		Conn: rc,
	}
	if err := psc.Subscribe(key); err != nil {
		return err
	}
	// get message from publisher
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return v
		}
	}
}
