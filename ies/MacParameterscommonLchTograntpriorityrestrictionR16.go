package ies

import "rrc/utils"

// MAC-ParametersCommon-lch-ToGrantPriorityRestriction-r16 ::= ENUMERATED
type MacParameterscommonLchTograntpriorityrestrictionR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLchTograntpriorityrestrictionR16EnumeratedNothing = iota
	MacParameterscommonLchTograntpriorityrestrictionR16EnumeratedSupported
)
