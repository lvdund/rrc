package ies

import "rrc/utils"

// UplinkConfig-enableDefaultBeamPL-ForPUSCH0-0-r16 ::= ENUMERATED
type UplinkconfigEnabledefaultbeamplForpusch00R16 struct {
	Value utils.ENUMERATED
}

const (
	UplinkconfigEnabledefaultbeamplForpusch00R16EnumeratedNothing = iota
	UplinkconfigEnabledefaultbeamplForpusch00R16EnumeratedEnabled
)
