package ies

import "rrc/utils"

// PUCCH-Config-dmrs-UplinkTransformPrecodingPUCCH-r16 ::= ENUMERATED
type PucchConfigDmrsUplinktransformprecodingpucchR16 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigDmrsUplinktransformprecodingpucchR16EnumeratedNothing = iota
	PucchConfigDmrsUplinktransformprecodingpucchR16EnumeratedEnabled
)
