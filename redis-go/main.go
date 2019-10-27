package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp",
		"XXXXX.c8XXXXXX-1-4.ec2.cloud.redislabs.com:XXXXXX",
		redis.DialPassword("0OhEKCh7XLllhES2aO1jU79EEvdRNWRf"))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	if _, err = conn.Do("SET", "key", "Hello World"); err != nil {
		log.Fatal(err)
	}

	str, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
