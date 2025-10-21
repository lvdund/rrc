package ies

import "rrc/utils"

// V2X-BandwidthClassSL-r14 ::= SEQUENCE OF V2X-BandwidthClass-r14
// SIZE (1..maxBandwidthClass-r10)
type V2xBandwidthclassslR14 struct {
	Value utils.Sequence[V2xBandwidthclassR14]
}
