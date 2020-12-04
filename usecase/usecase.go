package usecase

import (
	"context"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/model"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/service"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/umb"
)

type IUsecase interface {
	GetUniversalQuota(ctx context.Context, request *model.GetUniversalKuotaRq) (response *model.GetUniversalKuotaRs, err error)
}

type usecase struct {
	soarService service.ISoarService
}

func NewUseCases(soarService service.ISoarService) IUsecase {
	return &usecase{soarService: soarService}
}

type GetUniversalKuotaInput struct {
	ctx context.Context
	msisdn string
}

func (u * usecase) getUniversalKuota(in *GetUniversalKuotaInput) umb.UMB {
	srvPacketRq := new(model.GetUniversalKuotaRq)
	srvPacketRq.Msisdn = in.msisdn
	packetSrv, err := u.soarService.GetUniversalQuota(in.ctx, srvPacketRq)

}
