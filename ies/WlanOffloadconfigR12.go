package ies

import "rrc/utils"

// WLAN-OffloadConfig-r12 ::= SEQUENCE
// Extensible
type WlanOffloadconfigR12 struct {
	ThresholdrsrpR12                   *WlanOffloadconfigR12ThresholdrsrpR12
	ThresholdrsrqR12                   *WlanOffloadconfigR12ThresholdrsrqR12
	ThresholdrsrqOnallsymbolswithwbR12 *WlanOffloadconfigR12ThresholdrsrqOnallsymbolswithwbR12
	ThresholdrsrqOnallsymbolsR12       *WlanOffloadconfigR12ThresholdrsrqOnallsymbolsR12
	ThresholdrsrqWbR12                 *WlanOffloadconfigR12ThresholdrsrqWbR12
	ThresholdchannelutilizationR12     *WlanOffloadconfigR12ThresholdchannelutilizationR12
	ThresholdbackhaulBandwidthR12      *WlanOffloadconfigR12ThresholdbackhaulBandwidthR12
	ThresholdwlanRssiR12               *WlanOffloadconfigR12ThresholdwlanRssiR12
	OffloadpreferenceindicatorR12      *utils.BITSTRING `lb:16,ub:16`
	TSteeringwlanR12                   *TReselection
}
