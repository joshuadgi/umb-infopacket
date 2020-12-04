package umb

import "encoding/xml"

type UMB struct {
	XMLName xml.Name `xml:"http://schemas.xl.co.id/UMBResponse/V1.0 umb"`
	Menu UMBMenu
	Result UMBResult
}

type UMBMenu struct {
	XMLName xml.Name `xml:"menu"`
	TariffType string `xml:"tarifftype,attr"`
	TariffRate string `xml:"tariffrate,attr"`
	MenuHeader1 string `xml:"menuheader1"`
	MenuHeader2 string `xml:"menuheader2"`
	MenuName string `xml:"menuname"`
	Item []UMBMenuItem
}

type UMBMenuItem struct {
	XMLName xml.Name `xml:"item"`
	Number int `xml:"number,attr"`
	Action string `xml:"action,attr"`
	Url string `xml:"url,attr"`
	DeliveryMode string `xml:"deliverymode,attr"`
	TariffType string `xml:"tarifftype,attr"`
	TariffRate int `xml:"tariffrate,attr"`
	Value string `xml:"value"`
}

type UMBResult struct {
	XMLName xml.Name `xml:"result"`
	ResultData string `xml:"resultdata"`
}
