package ies

import "rrc/utils"

// SRS-ConfigAdd-r16-srs-BandwidthAdd-r16 ::= ENUMERATED
type SrsConfigaddR16SrsBandwidthaddR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigaddR16SrsBandwidthaddR16EnumeratedNothing = iota
	SrsConfigaddR16SrsBandwidthaddR16EnumeratedBw0
	SrsConfigaddR16SrsBandwidthaddR16EnumeratedBw1
	SrsConfigaddR16SrsBandwidthaddR16EnumeratedBw2
	SrsConfigaddR16SrsBandwidthaddR16EnumeratedBw3
)
