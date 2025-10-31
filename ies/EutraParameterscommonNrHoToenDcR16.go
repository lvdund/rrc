package ies

import "rrc/utils"

// EUTRA-ParametersCommon-nr-HO-ToEN-DC-r16 ::= ENUMERATED
type EutraParameterscommonNrHoToenDcR16 struct {
	Value utils.ENUMERATED
}

const (
	EutraParameterscommonNrHoToenDcR16EnumeratedNothing = iota
	EutraParameterscommonNrHoToenDcR16EnumeratedSupported
)
