package ies

import "rrc/utils"

// SRS-ConfigAdd-r16-srs-HoppingBandwidthAdd-r16 ::= ENUMERATED
type SrsConfigaddR16SrsHoppingbandwidthaddR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigaddR16SrsHoppingbandwidthaddR16EnumeratedNothing = iota
	SrsConfigaddR16SrsHoppingbandwidthaddR16EnumeratedHbw0
	SrsConfigaddR16SrsHoppingbandwidthaddR16EnumeratedHbw1
	SrsConfigaddR16SrsHoppingbandwidthaddR16EnumeratedHbw2
	SrsConfigaddR16SrsHoppingbandwidthaddR16EnumeratedHbw3
)
