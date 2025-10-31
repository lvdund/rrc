package ies

import "rrc/utils"

// MBMS-ROM-Info-r15-mbms-Bandwidth-r15 ::= ENUMERATED
type MbmsRomInfoR15MbmsBandwidthR15 struct {
	Value utils.ENUMERATED
}

const (
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedNothing = iota
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN6
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN15
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN25
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN50
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN75
	MbmsRomInfoR15MbmsBandwidthR15EnumeratedN100
)
