package ies

import "rrc/utils"

// SL-InterFreqInfoV2X-r14-sl-Bandwidth-r14 ::= ENUMERATED
type SlInterfreqinfov2xR14SlBandwidthR14 struct {
	Value utils.ENUMERATED
}

const (
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedNothing = iota
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN6
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN15
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN25
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN50
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN75
	SlInterfreqinfov2xR14SlBandwidthR14EnumeratedN100
)
