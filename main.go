package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func ExampleClient(client *redis.Client) {
	key := "visitors"
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("After Updating the vlaue of key is ", val)
	err = client.Set("visitors", val+1, 0).Err()
	if err != nil {
		panic(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis-server:6379",
	})
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	err := client.Set("visitors", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	ExampleClient(client)
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
