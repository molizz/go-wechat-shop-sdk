package shop

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productJSON = `
{
	"title": "眼镜大佬 - 微信小程序 - 在线配眼镜",
	"out_product_id": "hello-world",
	"path": "pages",
	"head_img": [
	   "https://myqcloud.com/product/zeiss-a-series.jpg"
	],
	"desc_info": {
	   "desc": "蔡司入门级镜片\n",
	   "imgs": [
		  "https://myqcloud.com/product/zeiss-a-series.jpg"
	   ]
	},
	"third_cat_id": 6112,
	"brand_id": 2100000000,
	"info_version": "1",
	"skus": [
	   {
		  "out_product_id": "zeiss-a-series",
		  "out_sku_id": "zeiss-a-series-1.56",
		  "thumb_img": "https://myqcloud.com/product/zeiss-a-series.jpg",
		  "sale_price": 280,
		  "market_price": 880,
		  "stock_num": 9999,
		  "sku_attrs": [
			 {
				"attr_key": "折射率",
				"attr_value": "1.56"
			 }
		  ]
	   },
	   {
		  "out_product_id": "zeiss-a-series",
		  "out_sku_id": "zeiss-a-series-1.60",
		  "thumb_img": "https://myqcloud.com/product/zeiss-a-series.jpg",
		  "sale_price": 380,
		  "market_price": 980,
		  "stock_num": 9999,
		  "sku_attrs": [
			 {
				"attr_key": "折射率",
				"attr_value": "1.6"
			 }
		  ]
	   },
	   {
		  "out_product_id": "zeiss-a-series",
		  "out_sku_id": "zeiss-a-series-1.67",
		  "thumb_img": "https://myqcloud.com/product/zeiss-a-series.jpg",
		  "sale_price": 480,
		  "market_price": 1580,
		  "stock_num": 9999,
		  "sku_attrs": [
			 {
				"attr_key": "折射率",
				"attr_value": "1.67"
			 }
		  ]
	   },
	   {
		  "out_product_id": "zeiss-a-series",
		  "out_sku_id": "zeiss-a-series-1.74",
		  "thumb_img": "https://myqcloud.com/product/zeiss-a-series.jpg",
		  "sale_price": 1280,
		  "market_price": 3880,
		  "stock_num": 9999,
		  "sku_attrs": [
			 {
				"attr_key": "折射率",
				"attr_value": "1.74"
			 }
		  ]
	   }
	]
 }
`

func TestSpu_Add(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	spu := NewSpu(at)
	product := new(SpuProduct)
	err = json.Unmarshal([]byte(productJSON), &product)
	assert.Nil(t, err)
	assert.Equal(t, "hello-world", product.OutProductID)

	result, err := spu.Add(product)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	err = spu.DelProduct(product.OutProductID)
	assert.Nil(t, err)
}

func TestSpu_Get(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	spu := NewSpu(at)
	product := new(SpuProduct)
	err = json.Unmarshal([]byte(productJSON), &product)
	assert.Nil(t, err)
	assert.Equal(t, "hello-world", product.OutProductID)

	result, err := spu.Add(product)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	defer func() {
		err = spu.DelProduct(result.Data.OutProductID)
		assert.Nil(t, err)
	}()

	getproduct, err := spu.Get(result.Data.OutProductID, 1)
	assert.Nil(t, err)
	assert.Equal(t, getproduct.Spu.OutProductID, result.Data.OutProductID)
}
