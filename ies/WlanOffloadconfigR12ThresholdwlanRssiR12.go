package ies

import "rrc/utils"

// WLAN-OffloadConfig-r12-thresholdWLAN-RSSI-r12 ::= SEQUENCE
type WlanOffloadconfigR12ThresholdwlanRssiR12 struct {
	ThresholdwlanRssiLowR12  utils.INTEGER `lb:0,ub:255`
	ThresholdwlanRssiHighR12 utils.INTEGER `lb:0,ub:255`
}
