package v3wechat

import "sync"

type Option func(c *Client)

func UpdateMchInfoOptions(f func(c *Client)) Option {
	return func(c *Client) {
		c.updateMchInfoFunc = f
	}
}

type Client struct {
	mu sync.Mutex
	mi MchInfo

	updateMchInfoFunc func(c *Client)
}

func NewClient(mi MchInfo, options ...Option) *Client {
	c := &Client{mi: mi}

	for _, o := range options {
		o(c)
	}
	return c
}

func (c *Client) GetMchInfo() MchInfo {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.mi
}

func (c *Client) UpdateMchInfo(mi MchInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mi = mi
}
