package ies

import "rrc/utils"

// MeasResultWLAN-r13 ::= SEQUENCE
// Extensible
type MeasresultwlanR13 struct {
	WlanIdentifiersR13                WlanIdentifiersR12
	CarrierinfowlanR13                *WlanCarrierinfoR13
	BandwlanR13                       *WlanBandindicatorR13
	RssiwlanR13                       WlanRssiRangeR13
	AvailableadmissioncapacitywlanR13 *utils.INTEGER `lb:0,ub:31250`
	BackhauldlBandwidthwlanR13        *WlanBackhaulrateR12
	BackhaululBandwidthwlanR13        *WlanBackhaulrateR12
	ChannelutilizationwlanR13         *utils.INTEGER `lb:0,ub:255`
	StationcountwlanR13               *utils.INTEGER `lb:0,ub:65535`
	ConnectedwlanR13                  *bool
}
