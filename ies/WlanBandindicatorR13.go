package ies

import "rrc/utils"

// WLAN-BandIndicator-r13 ::= ENUMERATED
// Extensible
type WlanBandindicatorR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanBandindicatorR13Band2dot4   = 0
	WlanBandindicatorR13Band5       = 1
	WlanBandindicatorR13Band60V1430 = 2
	WlanBandindicatorR13Spare5      = 3
	WlanBandindicatorR13Spare4      = 4
	WlanBandindicatorR13Spare3      = 5
	WlanBandindicatorR13Spare2      = 6
	WlanBandindicatorR13Spare1      = 7
)
