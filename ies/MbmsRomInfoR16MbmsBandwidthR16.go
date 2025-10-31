package ies

import "rrc/utils"

// MBMS-ROM-Info-r16-mbms-Bandwidth-r16 ::= ENUMERATED
type MbmsRomInfoR16MbmsBandwidthR16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedNothing = iota
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN6
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN15
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN25
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN50
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN75
	MbmsRomInfoR16MbmsBandwidthR16EnumeratedN100
)
