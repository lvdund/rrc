package ies

import "rrc/utils"

// WLAN-OffloadConfig-r12-thresholdChannelUtilization-r12 ::= SEQUENCE
type WlanOffloadconfigR12ThresholdchannelutilizationR12 struct {
	ThresholdchannelutilizationlowR12  utils.INTEGER `lb:0,ub:255`
	ThresholdchannelutilizationhighR12 utils.INTEGER `lb:0,ub:255`
}
