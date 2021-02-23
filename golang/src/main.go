package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/redis.v4"
)

//RedisClient function sets key in Redis to value provided in env var from docker-compose.yaml,
//pulls that newly set value and returns it as a string
func RedisClient() string {

	//bring in env vars from docker-compose file
	var RedisHost string
	RedisHost = os.Getenv("REDIS_HOST")
	var RedisPort string
	RedisPort = os.Getenv("REDIS_PORT")
	var DBstring string
	DBstring = RedisHost + ":" + RedisPort
	var ContentString string
	ContentString = os.Getenv("CONTENT_STRING")

	//connect to redis using values from env vars
	client := redis.NewClient(&redis.Options{
		Addr:     DBstring,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// using the value from env vars, set the key in Redis
	err := client.Set("my_content_string", ContentString, 0).Err()
	if err != nil {
		panic(err)
	}

	// read the previous value set in Redis
	val, err := client.Get("my_content_string").Result()
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	// Set var that will display as HTML content
	var ContentString string
	ContentString = RedisClient()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Value in Redis: %q", ContentString)
	})

	log.Fatal(http.ListenAndServe(":9001", nil))

}
