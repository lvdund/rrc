package ies

import "rrc/utils"

// SL-PreconfigGeneral-r12-sl-bandwidth-r12 ::= ENUMERATED
type SlPreconfiggeneralR12SlBandwidthR12 struct {
	Value utils.ENUMERATED
}

const (
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedNothing = iota
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN6
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN15
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN25
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN50
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN75
	SlPreconfiggeneralR12SlBandwidthR12EnumeratedN100
)
