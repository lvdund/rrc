package ies

import "rrc/utils"

// DL-CarrierConfigCommon-NB-r14-inbandCarrierInfo-r14-samePCI-Indicator-r14-samePCI-r14 ::= SEQUENCE
type DlCarrierconfigcommonNbR14InbandcarrierinfoR14SamepciIndicatorR14SamepciR14 struct {
	IndextomidprbR14 utils.INTEGER `lb:0,ub:54`
}
