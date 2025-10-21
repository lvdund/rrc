package ies

import "rrc/utils"

// UplinkPowerControlCommon-v1020 ::= SEQUENCE
type UplinkpowercontrolcommonV1020 struct {
	DeltafPucchFormat3R10    utils.ENUMERATED
	DeltafPucchFormat1bcsR10 utils.ENUMERATED
}
