package main

import (
	"github.com/labstack/echo/v4"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/delivery"
)

func main(){
	e:= echo.New()
	delivery.RestAPI(e)
	e.Logger.Fatal(e.Start("80"))
}

/*test := &umb.UMB{
Menu:    umb.UMBMenu{
TariffType:  "NONE",
TariffRate:  "0",
MenuHeader1: "menu1",
MenuName:    "Main Menu",
Item:        []umb.UMBMenuItem{{
Number:       1,
Action:       "NEXT",
Url:          "localhost:8080/nextTestMenu/test",
DeliveryMode: "0",
TariffType:   "NONE",
TariffRate:   0,
Value:        "Go to test",
},
},
},
}
out,_ := xml.MarshalIndent(test, " ", "  ")
fmt.Println(string(out))

fmt.Println(xml.Header + string(out))*/