package ies

import "rrc/utils"

// UplinkPowerControlCommonPSCell-r12 ::= SEQUENCE
type UplinkpowercontrolcommonpscellR12 struct {
	DeltafPucchFormat3R12    utils.ENUMERATED
	DeltafPucchFormat1bcsR12 utils.ENUMERATED
	P0NominalpucchR12        utils.INTEGER
	DeltaflistPucchR12       DeltaflistPucch
}
