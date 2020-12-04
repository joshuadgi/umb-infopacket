package delivery

import (
	"context"
	"encoding/xml"
	_ "encoding/xml"
	"fmt"
	"github.com/labstack/echo/v4"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/model"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/umb"
)

func (r *RestHandler) InfoPacket(c echo.Context) error {
	apiResponse := new(umb.UMB)
	request := new(model.GetUniversalKuotaRq)
	request.Msisdn = "62817100703"
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//err := c.Bind(request)
	usc, err := r.Usecase.GetUniversalKuota(ctx, request)
	if err != nil {
		//error handling
	}
	if usc != nil {
		items := []umb.UMBMenuItem{}
		for i, packet := range usc.ListPacket {
			items = append(items, umb.UMBMenuItem{
				Number:       i,
				Action:       "NEXT",
				Url:          fmt.Sprintf("localhost:8080/packet/%s/%d", packet.PType, packet.PId),
				DeliveryMode: "0",
				TariffType:   "NONE",
				TariffRate:   0,
				Value:        packet.PName,
			})
		}
		apiResponse.Menu = umb.UMBMenu{
			TariffType:  "NONE",
			TariffRate:  "0",
			MenuHeader1: "Welcome",
			MenuName:    "Menu Info Packet",
			Item:        items,
		}

	}
	xmlResponse, err := xml.Marshal(apiResponse)
	if err!=nil{
		//error handling
	}
	return respUMB(c,200, xmlResponse, err)
}
