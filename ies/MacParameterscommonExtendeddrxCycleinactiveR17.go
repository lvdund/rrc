package ies

import "rrc/utils"

// MAC-ParametersCommon-extendedDRX-CycleInactive-r17 ::= ENUMERATED
type MacParameterscommonExtendeddrxCycleinactiveR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonExtendeddrxCycleinactiveR17EnumeratedNothing = iota
	MacParameterscommonExtendeddrxCycleinactiveR17EnumeratedSupported
)
