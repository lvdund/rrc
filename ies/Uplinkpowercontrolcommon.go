package ies

import "rrc/utils"

// UplinkPowerControlCommon ::= SEQUENCE
type Uplinkpowercontrolcommon struct {
	P0Nominalpusch    utils.INTEGER `lb:0,ub:24`
	Alpha             AlphaR12
	P0Nominalpucch    utils.INTEGER `lb:0,ub:-96`
	DeltaflistPucch   DeltaflistPucch
	Deltapreamblemsg3 utils.INTEGER `lb:0,ub:6`
}
