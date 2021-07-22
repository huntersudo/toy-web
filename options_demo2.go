package main

type Client interface {
	Conn()
}

type DBClient struct {
	timeout    int
	retryTimes int
}

func (c DBClient) Conn() {
	// Do SOMETHING
}

type ClientConnOption func(*ClientConnOptions)

type ClientConnOptions struct {
	retryTimes int
	timeout    int
}

func WithRetryTimes(retryTimes int) ClientConnOption {
	return func(options *ClientConnOptions) {
		options.retryTimes = retryTimes
	}
}

// 闭包
func WithTimeout(timeout int) ClientConnOption {
	return func(options *ClientConnOptions) {
		options.timeout = timeout
	}
}

func NewClient(opts ...ClientConnOption) Client {
	var defaultClientOptions = ClientConnOptions{
		retryTimes: 3,
		timeout:    5,
	}
	options := defaultClientOptions

	for _, optFunc := range opts {
		optFunc(&options)
	}

	return &DBClient{
		timeout:    options.timeout,
		retryTimes: options.retryTimes,
	}
}

