package web

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Hook 是一个钩子函数。注意，
// ctx 是一个有超时机制的 context.Context
// 所以你必须处理超时的问题
type Hook func(ctx context.Context) error

// BuildCloseServerHook 这里其实可以考虑使用 errgroup，
// 但是我们这里不用是希望每个 server 单独关闭
// 互相之间不影响
// todo  闭包  js类似，java不习惯 ，很多这种写法
func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{})
		//  + n
		wg.Add(len(servers))

		for _, s := range servers {
			// 每一个server 都是 单独去关停
			go func(svr Server) {
				err := svr.Shutdown(ctx)
				if err != nil {
					fmt.Printf("server shutdown error: %v \n", err)
				}
				time.Sleep(time.Second)
				wg.Done()  //  -1
			}(s)
		}

		go func() {
			wg.Wait() //都关完 后，继续下一步
			doneCh <- struct{}{}
		}()
		// todo 这里考虑超时，所以引入 doneCh ,同时开一个goroutine 去wait，避免影响下面
		select {
		case <- ctx.Done(): // 返回一个ch，
			fmt.Printf("closing servers timeout \n")
			return ErrorHookTimeout
		case <- doneCh:
			fmt.Printf("close all servers \n")
			return nil
		}
	}
}
