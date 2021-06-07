package shop

const (
	baseAPI = "https://api.weixin.qq.com"
)

type Shop struct {
	accessToken string
}

func NewShop(accessToken string) *Shop {
	return &Shop{accessToken: accessToken}
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
