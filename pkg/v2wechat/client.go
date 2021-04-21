package v2wechat

type Option func(c *Client)

func GetMchInfoOptions(f func(c *Client)*MchInfo) Option {
	return func(c *Client) {
		c.getMchInfoFunc = f
	}
}

type Client struct {
	mchId string

	getMchInfoFunc func(c *Client) *MchInfo
}

func NewClient(mchId string, options... Option) *Client{
	c := &Client{mchId: mchId}


	for _, o := range options {
		o(c)
	}
	return c
}