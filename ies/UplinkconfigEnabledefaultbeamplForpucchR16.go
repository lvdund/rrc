package ies

import "rrc/utils"

// UplinkConfig-enableDefaultBeamPL-ForPUCCH-r16 ::= ENUMERATED
type UplinkconfigEnabledefaultbeamplForpucchR16 struct {
	Value utils.ENUMERATED
}

const (
	UplinkconfigEnabledefaultbeamplForpucchR16EnumeratedNothing = iota
	UplinkconfigEnabledefaultbeamplForpucchR16EnumeratedEnabled
)
