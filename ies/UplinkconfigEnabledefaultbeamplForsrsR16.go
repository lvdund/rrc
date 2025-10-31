package ies

import "rrc/utils"

// UplinkConfig-enableDefaultBeamPL-ForSRS-r16 ::= ENUMERATED
type UplinkconfigEnabledefaultbeamplForsrsR16 struct {
	Value utils.ENUMERATED
}

const (
	UplinkconfigEnabledefaultbeamplForsrsR16EnumeratedNothing = iota
	UplinkconfigEnabledefaultbeamplForsrsR16EnumeratedEnabled
)
