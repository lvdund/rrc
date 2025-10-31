package ies

import "rrc/utils"

// WLAN-RTT-r16-rttUnits-r16 ::= utils.ENUMERATED // Extensible
type WlanRttR16RttunitsR16 struct {
	Value utils.ENUMERATED
}

const (
	WlanRttR16RttunitsR16EnumeratedNothing = iota
	WlanRttR16RttunitsR16EnumeratedMicroseconds
	WlanRttR16RttunitsR16EnumeratedHundredsofnanoseconds
	WlanRttR16RttunitsR16EnumeratedTensofnanoseconds
	WlanRttR16RttunitsR16EnumeratedNanoseconds
	WlanRttR16RttunitsR16EnumeratedTenthsofnanoseconds
)
