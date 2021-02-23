package main

import (
	"fmt"
	"os"

	redis "gopkg.in/redis.v4"
)

func main() {
	//bring in env vars from docker-compose file
	var RedisHost string
	RedisHost = os.Getenv("REDIS_HOST")
	var RedisPort string
	RedisPort = os.Getenv("REDIS_PORT")
	var DBstring string
	DBstring = RedisHost + ":" + RedisPort

	//connect to redis using values from env vars
	client := redis.NewClient(&redis.Options{
		Addr:     DBstring,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//collect DB ping results
	pong, err := client.Ping().Result()
	// if ping does not get a proper response then fail and notify
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//otherwise report success
	fmt.Println(pong)

}
