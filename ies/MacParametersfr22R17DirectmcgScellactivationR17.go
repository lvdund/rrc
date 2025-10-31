package ies

import "rrc/utils"

// MAC-ParametersFR2-2-r17-directMCG-SCellActivation-r17 ::= ENUMERATED
type MacParametersfr22R17DirectmcgScellactivationR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfr22R17DirectmcgScellactivationR17EnumeratedNothing = iota
	MacParametersfr22R17DirectmcgScellactivationR17EnumeratedSupported
)
