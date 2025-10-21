package ies

import "rrc/utils"

// UplinkPowerControlCommonSCell-v1310 ::= SEQUENCE
type UplinkpowercontrolcommonscellV1310 struct {
	P0Nominalpucch           utils.INTEGER
	DeltaflistPucch          DeltaflistPucch
	DeltafPucchFormat3R12    *utils.ENUMERATED
	DeltafPucchFormat1bcsR12 *utils.ENUMERATED
	DeltafPucchFormat4R13    *utils.ENUMERATED
	DeltafPucchFormat513     *utils.ENUMERATED
}
