package ies

import "rrc/utils"

// WLAN-OffloadConfig-r12 ::= SEQUENCE
// Extensible
type WlanOffloadconfigR12 struct {
	ThresholdrsrpR12                   *interface{}
	ThresholdrsrqR12                   *interface{}
	ThresholdrsrqOnallsymbolswithwbR12 *interface{}
	ThresholdrsrqOnallsymbolsR12       *interface{}
	ThresholdrsrqWbR12                 *interface{}
	ThresholdchannelutilizationR12     *interface{}
	ThresholdbackhaulBandwidthR12      *interface{}
	ThresholdwlanRssiR12               *interface{}
	OffloadpreferenceindicatorR12      *utils.BITSTRING
	TSteeringwlanR12                   *TReselection
}
