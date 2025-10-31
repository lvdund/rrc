package ies

import "rrc/utils"

// SSB-Configuration-r16-halfFrameIndex-r16 ::= ENUMERATED
type SsbConfigurationR16HalfframeindexR16 struct {
	Value utils.ENUMERATED
}

const (
	SsbConfigurationR16HalfframeindexR16EnumeratedNothing = iota
	SsbConfigurationR16HalfframeindexR16EnumeratedZero
	SsbConfigurationR16HalfframeindexR16EnumeratedOne
)
