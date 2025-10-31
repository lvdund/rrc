package ies

import "rrc/utils"

// CrossCarrierSchedulingConfig-enableDefaultBeamForCCS-r16 ::= ENUMERATED
type CrosscarrierschedulingconfigEnabledefaultbeamforccsR16 struct {
	Value utils.ENUMERATED
}

const (
	CrosscarrierschedulingconfigEnabledefaultbeamforccsR16EnumeratedNothing = iota
	CrosscarrierschedulingconfigEnabledefaultbeamforccsR16EnumeratedEnabled
)
