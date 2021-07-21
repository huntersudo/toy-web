package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	WithTimeout()
	WithCancel()
	WithDeadline()
	WithValue()
}

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	// 输出2，说明在 ctx.Done()这里阻塞了两秒
	fmt.Println(end-start)
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<- ctx.Done()
		fmt.Println("context was canceled")
	}()
	// 确保我们的 goroutine进去执行了
	time.Sleep(time.Second)
	cancel()
	// 确保后面那句打印出来了
	time.Sleep(time.Second)
}

func WithDeadline() {
	// 设置两秒后超时   todo  链路超时控制
	ctx, cancel := context.WithDeadline(context.Background(),
		time.Now().Add(2 * time.Second))
	defer cancel()

	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	// 输出2，说明在 ctx.Done()这里阻塞了两秒
	fmt.Println(end-start)
}

func WithValue() {
	parentKey := "parent"
	parent := context.WithValue(context.Background(), parentKey, "this is parent")

	sonKey := "son"
	son := context.WithValue(parent, sonKey, "this is son")
     // todo son可以拿到parent的，反之不行
	// 尝试从 parent 里面拿出来 key = son的，会拿不到
	if parent.Value(parentKey) == nil {
		fmt.Printf("parent can not get son's key-value pair")
	}

	if val := son.Value(parentKey); val != nil {
		fmt.Printf("parent can not get son's key-value pair")
	}

	//todo thread -local 线程安全的保存数据的地方
	// thread1 -> map[string]value
	// thread2 -> map[string]value
	//  互相不能访问
	 // Go 要么用channel，要么用 context
}
