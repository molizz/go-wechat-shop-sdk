package shop

import (
	"fmt"
)

type Spu struct {
	accessToken string
}

func NewSpu(accessToken string) *Spu {
	return &Spu{accessToken: accessToken}
}

type SpuAddResult struct {
	Result
	Data struct {
		ProductID    int64  `json:"product_id"`
		OutProductID string `json:"out_product_id"`
		CreateTime   string `json:"create_time"`
		Skus         []struct {
			SkuID    int    `json:"sku_id"`
			OutSkuID string `json:"out_sku_id"`
		} `json:"skus"`
	} `json:"data"`
}

type SpuProduct struct {
	OutProductID      string   `json:"out_product_id"`
	Title             string   `json:"title"`
	Path              string   `json:"path"`
	HeadImg           []string `json:"head_img"`
	QualificationPics []string `json:"qualification_pics,omitempty"`
	DescInfo          *struct {
		Desc *string  `json:"desc,omitempty"`
		Imgs []string `json:"imgs,omitempty"`
	} `json:"desc_info,omitempty"`
	ThirdCatID  int     `json:"third_cat_id"`
	BrandID     int     `json:"brand_id"`
	InfoVersion *string `json:"info_version,omitempty"`
	Skus        []struct {
		OutProductID string  `json:"out_product_id"`
		OutSkuID     string  `json:"out_sku_id"`
		ThumbImg     string  `json:"thumb_img"`
		SalePrice    int     `json:"sale_price"`
		MarketPrice  int     `json:"market_price"`
		StockNum     int     `json:"stock_num"`
		SkuCode      *string `json:"sku_code,omitempty"`
		Barcode      *string `json:"barcode,omitempty"`
		SkuAttrs     []struct {
			AttrKey   string `json:"attr_key"`
			AttrValue string `json:"attr_value"`
		} `json:"sku_attrs"`
	} `json:"skus"`
}

/*
Add 添加商品
新增成功后会直接提交审核，可通过商品审核回调，或者通过get接口的edit_status查看是否通过审核。
商品有2份数据，草稿和线上数据。
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/SPU/add_spu.html
*/
func (s *Spu) Add(product *SpuProduct) (*SpuAddResult, error) {
	result := new(SpuAddResult)
	_, err := POST(s.accessToken, "shop/spu/add", product, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}

/*
Update 更新商品
注意：更新成功后会更新到草稿数据并直接提交审核，审核完成后有回调，也可通过get接口的edit_status查看是否通过审核。

商品有两份数据，草稿和线上数据。

调用接口新增或修改商品数据后，影响的只是草稿数据，审核通过草稿数据才会覆盖线上数据正式生效。

注意：

third_cat_id请根据获取类目接口拿到，并确定其qualification_type类目资质是否为必填，若为必填，那么要先调类目资质审核接口进行该third_cat_id的资质审核；
qualification_pics请根据获取类目接口中对应third_cat_id的product_qualification_type为依据，若为必填，那么该字段需要加上该商品的资质图片；
若需要上传某品牌商品，需要按照微信小商店开通规则开通对应品牌使用权限。微信小商店品牌开通规则：点击跳转，若无品牌可指定无品牌(无品牌brand_id: 2100000000)。

https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/SPU/add_spu.html
*/
func (s *Spu) Update(product *SpuProduct) (*SpuAddResult, error) {
	result := new(SpuAddResult)
	_, err := POST(s.accessToken, "shop/spu/update", product, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}

type SpuProductReq struct {
	OutProductID string `json:"out_product_id,omitempty"`
	NeedEditSpu  int    `json:"need_edit_spu,omitempty"`
}

/*
DelProduct 删除商品
从初始值/上架/若干下架状态转换成逻辑删除（删除后不可恢复）
*/
func (s *Spu) DelProduct(outProductID string) error {
	result := new(Result)
	from := &SpuProductReq{
		OutProductID: outProductID,
	}

	_, err := POST(s.accessToken, "shop/spu/del", from, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

/*
DelAudit 撤回商品审核
对于审核中（edit_status=2）的商品无法重复提交，需要调用此接口，使商品流转进入未审核的状态（edit_status=1）,即可重新提交商品。
*/
func (s *Spu) DelAudit(outProductID string) error {
	result := new(Result)
	from := &SpuProductReq{
		OutProductID: outProductID,
	}

	_, err := POST(s.accessToken, "shop/spu/del_audit", from, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

type GetSpuProduct struct {
	Result
	Spu *SpuProduct `json:"spu"`
}

/*
Get 获取商品
need_edit_spu  // 默认0:获取线上数据, 1:获取草稿数据
*/
func (s *Spu) Get(outProductID string, needEditSpu int) (*GetSpuProduct, error) {
	if needEditSpu != 0 && needEditSpu != 1 {
		return nil, fmt.Errorf("invalid param 'needEditSpu', expected 0 or 1, got '%d'", needEditSpu)
	}
	result := new(GetSpuProduct)
	from := &SpuProductReq{
		OutProductID: outProductID,
		NeedEditSpu:  needEditSpu,
	}
	_, err := POST(s.accessToken, "shop/spu/get", from, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}

type SpuProductsResult struct {
	Result
	TotalNum int           `json:"total_num"`
	Spus     []*SpuProduct `json:"spus"`
}

// Seek
// 循环所有的Product，遇到任何错误将返回
func (p *SpuProductsResult) Seek(f func(*SpuProduct) error) error {
	var err error
	for _, s := range p.Spus {
		err = f(s)
		if err != nil {
			break
		}
	}
	return err
}

type SpuProductGetList struct {
	Status          int    `json:"status"`            // 选填，不填时获取所有状态商品
	StartCreateTime string `json:"start_create_time"` // 选填，与end_create_time成对
	EndCreateTime   string `json:"end_create_time"`   // 选填，与start_create_time成对
	StartUpdateTime string `json:"start_update_time"` // 选填，与end_update_time成对
	EndUpdateTime   string `json:"end_update_time"`   // 选填，与start_update_time成对
	Page            int    `json:"page"`              //
	PageSize        int    `json:"page_size"`         // 不超过100
	NeedEditSpu     int    `json:"need_edit_spu"`     // 默认0:获取线上数据, 1:获取草稿数据
}

/*
GetList 获取商品列表
时间范围 create_time 和 update_time 同时存在时，以 create_time 的范围为准
*/
func (s *Spu) GetList(from *SpuProductGetList) (*SpuProductsResult, error) {
	result := new(SpuProductsResult)
	_, err := POST(s.accessToken, "shop/spu/get_list", from, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}

/*
Listing 上架商品
如果该商品处于自主下架状态，调用此接口可把直接把商品重新上架 该接口不影响已经在审核流程的草稿数据
*/
func (s *Spu) Listing(outProductID string) error {
	result := new(SpuProductsResult)
	from := &SpuProductReq{
		OutProductID: outProductID,
	}
	_, err := POST(s.accessToken, "shop/spu/listing", from, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

/*
Delisting 下架商品
如果该商品处于自主下架状态，调用此接口可把直接把商品重新上架 该接口不影响已经在审核流程的草稿数据
*/
func (s *Spu) Delisting(outProductID string) error {
	result := new(SpuProductsResult)
	from := &SpuProductReq{
		OutProductID: outProductID,
	}
	_, err := POST(s.accessToken, "shop/spu/delisting", from, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}
