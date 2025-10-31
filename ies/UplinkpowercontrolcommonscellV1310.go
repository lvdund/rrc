package ies

import "rrc/utils"

// UplinkPowerControlCommonSCell-v1310 ::= SEQUENCE
type UplinkpowercontrolcommonscellV1310 struct {
	P0Nominalpucch           utils.INTEGER `lb:0,ub:-96`
	DeltaflistPucch          DeltaflistPucch
	DeltafPucchFormat3R12    *UplinkpowercontrolcommonscellV1310DeltafPucchFormat3R12
	DeltafPucchFormat1bcsR12 *UplinkpowercontrolcommonscellV1310DeltafPucchFormat1bcsR12
	DeltafPucchFormat4R13    *UplinkpowercontrolcommonscellV1310DeltafPucchFormat4R13
	DeltafPucchFormat513     *UplinkpowercontrolcommonscellV1310DeltafPucchFormat513
}
