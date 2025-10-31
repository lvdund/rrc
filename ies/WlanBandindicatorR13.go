package ies

import "rrc/utils"

// WLAN-BandIndicator-r13 ::= utils.ENUMERATED // Extensible
type WlanBandindicatorR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanBandindicatorR13EnumeratedNothing = iota
	WlanBandindicatorR13EnumeratedBand2dot4
	WlanBandindicatorR13EnumeratedBand5
	WlanBandindicatorR13EnumeratedBand60_V1430
	WlanBandindicatorR13EnumeratedSpare5
	WlanBandindicatorR13EnumeratedSpare4
	WlanBandindicatorR13EnumeratedSpare3
	WlanBandindicatorR13EnumeratedSpare2
	WlanBandindicatorR13EnumeratedSpare1
)
