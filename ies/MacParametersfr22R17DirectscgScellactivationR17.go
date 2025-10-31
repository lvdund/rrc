package ies

import "rrc/utils"

// MAC-ParametersFR2-2-r17-directSCG-SCellActivation-r17 ::= ENUMERATED
type MacParametersfr22R17DirectscgScellactivationR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfr22R17DirectscgScellactivationR17EnumeratedNothing = iota
	MacParametersfr22R17DirectscgScellactivationR17EnumeratedSupported
)
