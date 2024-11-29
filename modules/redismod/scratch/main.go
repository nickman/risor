package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://:qwerty@localhost:6379/1?dial_timeout=5s")
	//opt, err := redis.ParseURL("unix://nicholas@/path/to/redis.sock?db=1")
	if err != nil {
		panic(err)
	}
	fmt.Println("addr is", opt.Addr)

	fmt.Println("db is", opt.DB)
	fmt.Println("password is", opt.Password)
	fmt.Println("dial timeout is", opt.DialTimeout)

	// Create client as usually.
	//_ = redis.NewClient(opt)

}
