package ies

import "rrc/utils"

// DL-CarrierConfigDedicated-NB-r13-inbandCarrierInfo-r13-samePCI-Indicator-r13-samePCI-r13 ::= SEQUENCE
type DlCarrierconfigdedicatedNbR13InbandcarrierinfoR13SamepciIndicatorR13SamepciR13 struct {
	IndextomidprbR13 utils.INTEGER `lb:0,ub:54`
}
