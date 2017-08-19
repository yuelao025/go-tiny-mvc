package models


type Sku struct {
	Goods_SN string
	Name string
}

func SkuInfo() (Sku)  {

	data := Sku{
		Goods_SN:"skuno11111",
		Name:"iphone6",
	}

	return data

}


