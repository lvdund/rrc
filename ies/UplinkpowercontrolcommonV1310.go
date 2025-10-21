package ies

import "rrc/utils"

// UplinkPowerControlCommon-v1310 ::= SEQUENCE
type UplinkpowercontrolcommonV1310 struct {
	DeltafPucchFormat4R13 *utils.ENUMERATED
	DeltafPucchFormat513  *utils.ENUMERATED
}
