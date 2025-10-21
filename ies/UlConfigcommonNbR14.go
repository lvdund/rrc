package ies

import "rrc/utils"

// UL-ConfigCommon-NB-r14 ::= SEQUENCE
// Extensible
type UlConfigcommonNbR14 struct {
	UlCarrierfreqR14        CarrierfreqNbR13
	NprachParameterslistR14 *NprachParameterslistNbR14
}
