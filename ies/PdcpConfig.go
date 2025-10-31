package ies

import "rrc/utils"

// PDCP-Config ::= SEQUENCE
// Extensible
type PdcpConfig struct {
	Drb                          *PdcpConfigDrb
	Morethanonerlc               *PdcpConfigMorethanonerlc
	TReordering                  *PdcpConfigTReordering
	Cipheringdisabled            *utils.BOOLEAN                                    `ext`
	DiscardtimerextR16           *utils.Setuprelease[DiscardtimerextR16]           `ext`
	MorethantworlcDrbR16         *PdcpConfigMorethantworlcDrbR16                   `ext`
	EthernetheadercompressionR16 *utils.Setuprelease[EthernetheadercompressionR16] `ext`
	SurvivaltimestatesupportR17  *utils.BOOLEAN                                    `ext`
	UplinkdatacompressionR17     *utils.Setuprelease[UplinkdatacompressionR17]     `ext`
	Discardtimerext2R17          *utils.Setuprelease[Discardtimerext2R17]          `ext`
	InitialrxDelivR17            *utils.BITSTRING                                  `lb:32,ub:32,ext`
}
