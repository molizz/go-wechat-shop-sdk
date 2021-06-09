package shop

type Img struct {
	accessToken string
}

func NewImg(accessToken string) *Img {
	return &Img{accessToken: accessToken}
}

type ImgResult struct {
	Result
	ImgInfo struct {
		MediaID string `json:"media_id"`
	} `json:"img_info"`
}

/*
Upload
此接口目前只用于品牌申请和类目申请。
*/
func (i *Img) Upload(uploadFilePath string) (*ImgResult, error) {
	var result = &ImgResult{}
	err := Upload(i.accessToken, "shop/img/upload", uploadFilePath, result)
	if err != nil {
		return nil, err
	}

	if !result.OK() {
		return nil, result
	}
	return result, nil
}
