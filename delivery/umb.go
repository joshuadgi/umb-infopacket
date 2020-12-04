package delivery

import (
	"github.com/labstack/echo/v4"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/service"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/usecase"
)

type RestHandler struct {
	Usecase usecase.IUsecase
}

func RestAPI(e *echo.Echo){
	soarService := service.NewSoarService()

	handler := &RestHandler{Usecase:usecase.NewUseCases(soarService)}

	//Routes
	apis := e.Group("api/v1")
	apis.Add("GET", "/info-packet", handler.InfoPacket)
}

func respUMB(c echo.Context, sc int, xmlblob []byte, err error) error {
	return c.XMLBlob(sc, xmlblob)
}
