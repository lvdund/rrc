package ies

import "rrc/utils"

// UplinkPowerControlCommonPSCell-r12 ::= SEQUENCE
type UplinkpowercontrolcommonpscellR12 struct {
	DeltafPucchFormat3R12    UplinkpowercontrolcommonpscellR12DeltafPucchFormat3R12
	DeltafPucchFormat1bcsR12 UplinkpowercontrolcommonpscellR12DeltafPucchFormat1bcsR12
	P0NominalpucchR12        utils.INTEGER `lb:0,ub:-96`
	DeltaflistPucchR12       DeltaflistPucch
}
