package shop

/*
错误码： https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/errorcode.html
*/

const (
	baseAPI = "https://api.weixin.qq.com"
)

type AccessTokenGetter interface {
	Get() (string, error)
}

type Shop struct {
	accessToken string
}

func NewShop(accessToken string) *Shop {
	return &Shop{accessToken: accessToken}
}

func NewShopFromAccessTokenGetter(atGetter AccessTokenGetter) *Shop {
	at, err := atGetter.Get()
	if err != nil {
		panic(err)
	}
	return NewShop(at)
}

func (s *Shop) Register() *Register {
	return NewRegister(s.accessToken)
}

func (s *Shop) Cat() *Cat {
	return NewCat(s.accessToken)
}

func (s *Shop) Img() *Img {
	return NewImg(s.accessToken)
}

func (s *Shop) Spu() *Spu {
	return NewSpu(s.accessToken)
}

func (s *Shop) Order() *Order {
	return NewOrder(s.accessToken)
}

func (s *Shop) Delivery() *Delivery {
	return NewDelivery(s.accessToken)
}

func (s *Shop) Aftersale() *Aftersale {
	return NewAftersale(s.accessToken)
}
