package main

import (
	"fmt"
	xl_umb "gitlab.intra.excelcom.co.id/go/libraries/xl-umb"
)

func main() {
	//sample create menu item
	testItem1 := xl_umb.ConstructItem(1, "NEXT", "localhost:8080/test1", "0", "NONE", 0, "Menu Test")
	testItem2 := xl_umb.ConstructItem(2, "NEXT", "localhost:8080/test2", "0", "NONE", 0, "Menu Test")
	//add to items
	testItems := []xl_umb.UMBMenuItem{testItem1, testItem2}
	//create umb menu
	testMenu := xl_umb.UMBMenu{
		TariffType:  "NONE",
		TariffRate:  0,
		MenuHeader1: "Menu header",//optional
		MenuHeader2: "",//optional
		MenuName:    "Menu sample",
		Item:        testItems,
	}
	//construct xml
	xmlUmb, err := xl_umb.ConstructMenu(&testMenu)
	if err != nil {
		fmt.Println("cannot render xml")
	}else{
		//output version for echo xmlblob, can be read by parsing to string
		fmt.Println("string version :", string(xmlUmb))
	}

	//sample create result
	testResult := xl_umb.UMBResult{
		ResultData: "Test output result",
	}
	//construct xml
	xmlResUmb, err := xl_umb.ConstructResult(&testResult)
	if err != nil {
		fmt.Println("string render xml")
	}else{
		//output version for echo xmlblob, can be read by parsing to string
		fmt.Println("byte version :", string(xmlResUmb))
	}

}

