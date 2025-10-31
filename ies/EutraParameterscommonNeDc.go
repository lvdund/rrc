package ies

import "rrc/utils"

// EUTRA-ParametersCommon-ne-DC ::= ENUMERATED
type EutraParameterscommonNeDc struct {
	Value utils.ENUMERATED
}

const (
	EutraParameterscommonNeDcEnumeratedNothing = iota
	EutraParameterscommonNeDcEnumeratedSupported
)
