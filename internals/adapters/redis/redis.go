package redis

import (
	"context"
	"log"
	goredis "github.com/redis/go-redis/v9"
)


func NewClient(addr string) *goredis.Client {	
   rdb := goredis.NewClient(&goredis.Options{Addr: addr})
   
   if err := rdb.Ping(context.Background()).Err(); err != nil {	
	 log.Fatalf("failed to connect to Redis: %v", err)
   }

   log.Printf("Connected to Redis at %s", addr)

   return rdb
}
