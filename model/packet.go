package model

type Packet struct {
	PId int `json:"id"`
	PType string `json:"type"`
	PName string `json:"name"`
}

type GetUniversalKuotaRs struct {
	ListPacket []Packet `json:"listPacket"`
}

type GetUniversalKuotaRq struct {
	Msisdn string `json:"msisdn"`
}
