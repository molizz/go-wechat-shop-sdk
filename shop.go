package shop

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
