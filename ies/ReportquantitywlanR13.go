package ies

import "rrc/utils"

// ReportQuantityWLAN-r13 ::= SEQUENCE
// Extensible
type ReportquantitywlanR13 struct {
	BandrequestwlanR13                       *utils.ENUMERATED
	CarrierinforequestwlanR13                *utils.ENUMERATED
	AvailableadmissioncapacityrequestwlanR13 *utils.ENUMERATED
	BackhauldlBandwidthrequestwlanR13        *utils.ENUMERATED
	BackhaululBandwidthrequestwlanR13        *utils.ENUMERATED
	ChannelutilizationrequestwlanR13         *utils.ENUMERATED
	StationcountrequestwlanR13               *utils.ENUMERATED
}
