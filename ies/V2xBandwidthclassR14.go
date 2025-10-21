package ies

import "rrc/utils"

// V2X-BandwidthClass-r14 ::= ENUMERATED
// Extensible
type V2xBandwidthclassR14 struct {
	Value utils.ENUMERATED
}

const (
	V2xBandwidthclassR14A       = 0
	V2xBandwidthclassR14B       = 1
	V2xBandwidthclassR14C       = 2
	V2xBandwidthclassR14D       = 3
	V2xBandwidthclassR14E       = 4
	V2xBandwidthclassR14F       = 5
	V2xBandwidthclassR14C1V1530 = 6
)
