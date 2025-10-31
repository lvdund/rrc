package ies

import "rrc/utils"

// WLAN-RTT-r15-rttUnits-r15 ::= utils.ENUMERATED // Extensible
type WlanRttR15RttunitsR15 struct {
	Value utils.ENUMERATED
}

const (
	WlanRttR15RttunitsR15EnumeratedNothing = iota
	WlanRttR15RttunitsR15EnumeratedMicroseconds
	WlanRttR15RttunitsR15EnumeratedHundredsofnanoseconds
	WlanRttR15RttunitsR15EnumeratedTensofnanoseconds
	WlanRttR15RttunitsR15EnumeratedNanoseconds
	WlanRttR15RttunitsR15EnumeratedTenthsofnanoseconds
)
