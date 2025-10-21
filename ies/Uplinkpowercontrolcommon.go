package ies

import "rrc/utils"

// UplinkPowerControlCommon ::= SEQUENCE
type Uplinkpowercontrolcommon struct {
	P0Nominalpusch    utils.INTEGER
	Alpha             AlphaR12
	P0Nominalpucch    utils.INTEGER
	DeltaflistPucch   DeltaflistPucch
	Deltapreamblemsg3 utils.INTEGER
}
