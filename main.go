package main

import (
  "context"
  "fmt"
  "time"

  "github.com/redis/go-redis/v9"
)

func main() {
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
  })

  ctx := context.Background()

  // Set key
  err := rdb.Set(ctx, "key", "value", 0).Err()
  if err != nil {
    panic(err)
  }

  // Get key
  val, err := rdb.Get(ctx, "key").Result()
  if err != nil {
    panic(err)
  }
  fmt.Println("key", val)

  // Set with expiry
  err = rdb.Set(ctx, "key2", "value2", 10*time.Second).Err()
  if err != nil {
    panic(err)
  }

  fmt.Println("Redis cache demo completed")
}