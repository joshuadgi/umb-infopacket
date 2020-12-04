package usecase

import (
	"context"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/model"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/service"
)

type IUsecase interface {
	GetUniversalKuota(ctx context.Context, request *model.GetUniversalKuotaRq) (response *model.GetUniversalKuotaRs, err error)
}

type usecase struct {
	soarService service.ISoarService
}

func NewUseCases(soarService service.ISoarService) IUsecase {
	return &usecase{soarService: soarService}
}

func (u * usecase) GetUniversalKuota(ctx context.Context, request *model.GetUniversalKuotaRq) (response *model.GetUniversalKuotaRs, err error) {
	srvPacketRq := new(model.GetUniversalKuotaRq)
	srvPacketRq.Msisdn = request.Msisdn
	packetSrv, err := u.soarService.GetUniversalQuota(ctx, srvPacketRq)
	if err!=nil {
		return nil, err
	}
	return packetSrv, nil
}
