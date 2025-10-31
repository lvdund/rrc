package ies

import "rrc/utils"

// V2X-BandwidthClass-r14 ::= utils.ENUMERATED // Extensible
type V2xBandwidthclassR14 struct {
	Value utils.ENUMERATED
}

const (
	V2xBandwidthclassR14EnumeratedNothing = iota
	V2xBandwidthclassR14EnumeratedA
	V2xBandwidthclassR14EnumeratedB
	V2xBandwidthclassR14EnumeratedC
	V2xBandwidthclassR14EnumeratedD
	V2xBandwidthclassR14EnumeratedE
	V2xBandwidthclassR14EnumeratedF
	V2xBandwidthclassR14EnumeratedC1_V1530
)
