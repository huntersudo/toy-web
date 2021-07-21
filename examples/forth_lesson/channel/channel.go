package main

import (
	"fmt"
	"time"
)

func main() {
	channelWithoutCache()
	channelWithCache()
	// Hello, msg from channel
	//2021-07-21 20:48:40.0760539 +0800 CST m=+3.007654001Hello, first msg from channel
	//2021-07-21 20:48:40.146777 +0800 CST m=+3.078377101Hello, second msg from channel
}

func channelWithCache()  {
	ch := make(chan string, 1)
	go func() {

		ch <- "Hello, first msg from channel"
		time.Sleep(time.Second)
		ch <- "Hello, second msg from channel"
	}()

	time.Sleep(2 * time.Second)
	msg := <- ch
	fmt.Println(time.Now().String() + msg)
	msg = <- ch
	fmt.Println(time.Now().String() + msg)
	// todo 因为前面我们先睡了2秒，所以其实会有一个已经在缓冲了
	// 当我们尝试输出的时候，这个输出间隔就会明显小于1秒
	// 我电脑上的几次实验，差距都在1ms以内
}

func channelWithoutCache() {
	// 不带缓冲
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "Hello, msg from channel"
	}()

	// 这里比较容易写成 msg <- ch，编译会报错
	msg := <- ch
	fmt.Println(msg)
}
