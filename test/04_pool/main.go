/*
 * @Author: v_wurunxin
 * @Date: 2022-08-04 15:19:45
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

type pool struct {
	work chan func()   // 任务
	sem  chan struct{} // 数量
}

func NewPool(num int) *pool {
	// work 无缓冲通道
	// sem 有缓冲通道，协程数
	return &pool{
		work: make(chan func()),
		sem:  make(chan struct{}, num),
	}
}

func (p *pool) NewTask(task func()) {
	// 第一次请求，无缓冲通道work不成立，进入开启协程逻辑
	select {
	case p.work <- task:
		fmt.Println("work")
		// 进入这里说明有协程执行完空闲阻塞着
	case p.sem <- struct{}{}:
		fmt.Println("sem")
		go p.worker(task)
	}
}

func (p *pool) worker(task func()) {
	defer func() { <-p.sem }()
	// 协程不关闭，等待继续执行
	for {
		task()
		// 阻塞
		fmt.Println("task end")
		task = <-p.work
	}
}

func main() {
	pool := NewPool(2)
	for i := 0; i <= 3; i++ {
		pool.NewTask(func() {
			fmt.Println(time.Now())
			time.Sleep(2 * time.Second)
		})
	}
	time.Sleep(5 * time.Second)
}
