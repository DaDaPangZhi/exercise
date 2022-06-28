/*
 * @Author: v_wurunxin
 * @Date: 2022-06-28 09:59:26
 * @Description: rate limit 测试
 */
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

/**
 * @description: wait测试，qps = 6 并发 = 3
 * @return {*}
 */
func wait() {
	limiter := rate.NewLimiter(6, 3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
	}
}

/**
 * @description: allow测试，qps = 4 并发 = 3
 * @return {*}
 */
func allow() {
	limiter := rate.NewLimiter(4, 1)
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("%03d Ok  %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			fmt.Printf("%03d Err %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// wait()
	allow()
}
