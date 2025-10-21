package ies

import "rrc/utils"

// UplinkPowerControlCommon-v1610 ::= SEQUENCE
type UplinkpowercontrolcommonV1610 struct {
	AlphasrsAddR16     AlphaR12
	P0NominalsrsAddR16 utils.INTEGER
}
