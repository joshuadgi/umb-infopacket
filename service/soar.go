package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitrepo.intra.excelcom.co.id/anthos/application-code/umb-infopacket/model"
	"time"
)

func NewSoarService() ISoarService {
return &soarService{rest: resty.New()}
}

type soarService struct {
	GetUniversalKuota *model.GetUniversalKuotaRq
	rest *resty.Client
}

type ISoarService interface {
	GetUniversalQuota(ctx context.Context, request *model.GetUniversalKuotaRq) (response *model.GetUniversalKuotaRs, err error)
}

func (s soarService) GetUniversalQuota (ctx context.Context, request *model.GetUniversalKuotaRq) (response *model.GetUniversalKuotaRs, err error) {
	rs := new(model.GetUniversalKuotaRs)
	url := "https://run.mocky.io/v3/d574b085-140b-43f4-a750-43e8794a5f7b"

	soarPayload := new(model.GetUniversalKuotaRq)
	soarPayload.Msisdn = request.Msisdn

	bodyRq, err := json.Marshal(soarPayload)
	if err != nil {
		return nil, fmt.Errorf("error parsing json : %s", err.Error())
	}

	//Resty Call HTTP
	client := s.rest
	client.SetTimeout(time.Duration(10000))
	soarRs, err := client.R().
		 SetHeader("Content-Type", "application/json").
		 SetBody(bodyRq).
		 SetContext(ctx).
		 SetResult(response).
		 Post(url)
	if err != nil {
		return nil, fmt.Errorf("error calling soar :", err.Error())
	}
	fmt.Println(soarRs)
	return response, err
}
