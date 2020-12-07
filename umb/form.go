package xl_umb

import "encoding/xml"

type UMB struct {
	XMLName xml.Name `xml:"ns0:umb"`
	XMLNs string `xml:"xmlns:ns0,attr"`
	Menu *UMBMenu
	Result *UMBResult
}

type UMBMenu struct {
	XMLName xml.Name `xml:"ns0:menu,omitempty"`
	TariffType string `xml:"tarifftype,attr"`
	TariffRate int `xml:"tariffrate,attr"`
	MenuHeader1 string `xml:"ns0:menuheader1,omitempty"`
	MenuHeader2 string `xml:"ns0:menuheader2,omitempty"`
	MenuName string `xml:"ns0:menuname"`
	Item []UMBMenuItem
}

type UMBMenuItem struct {
	XMLName xml.Name `xml:"ns0:item"`
	Number int `xml:"number,attr"`
	Action string `xml:"action,attr"`
	Url string `xml:"url,attr"`
	DeliveryMode string `xml:"deliverymode,attr"`
	TariffType string `xml:"tarifftype,attr"`
	TariffRate int `xml:"tariffrate,attr"`
	Value string `xml:",chardata"`
}

type UMBResult struct {
	XMLName xml.Name `xml:"ns0:result,omitempty"`
	ResultData string `xml:"ns0:resultdata"`
}
