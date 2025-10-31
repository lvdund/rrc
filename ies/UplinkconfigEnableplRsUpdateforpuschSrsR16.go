package ies

import "rrc/utils"

// UplinkConfig-enablePL-RS-UpdateForPUSCH-SRS-r16 ::= ENUMERATED
type UplinkconfigEnableplRsUpdateforpuschSrsR16 struct {
	Value utils.ENUMERATED
}

const (
	UplinkconfigEnableplRsUpdateforpuschSrsR16EnumeratedNothing = iota
	UplinkconfigEnableplRsUpdateforpuschSrsR16EnumeratedEnabled
)
