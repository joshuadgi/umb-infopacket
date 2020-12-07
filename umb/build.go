package xl_umb

import "encoding/xml"

func ConstructMenu (menu *UMBMenu) ([]byte, error){
	umb := UMB{
		XMLNs: "http://schemas.xl.co.id/UMBResponse/V1.0 ",
		Menu:    menu,
	}
	xmlUmb, err := xml.Marshal(umb)
	if err != nil {
		return nil, err
	}
	return xmlUmb, nil
}

func ConstructResult (result *UMBResult) ([]byte, error){
	umb :=  UMB{
		XMLNs: "http://schemas.xl.co.id/UMBResponse/V1.0 ",
		Result:  result,
	}
	xmlUmb, err := xml.Marshal(umb)
	if err != nil {
		return nil, err
	}
	return xmlUmb, nil
}

func ConstructItem (number int, action string, url string, deliverymode string, tarifftype string, tariffrate int, value string) (UMBMenuItem){
	return UMBMenuItem{
		Number:       number,
		Action:       action,
		Url:          url,
		DeliveryMode: deliverymode,
		TariffType:   tarifftype,
		TariffRate:   tariffrate,
		Value:        value,
	}
}
