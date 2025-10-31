package ies

// ReportQuantityWLAN-r13 ::= SEQUENCE
// Extensible
type ReportquantitywlanR13 struct {
	BandrequestwlanR13                       *bool
	CarrierinforequestwlanR13                *bool
	AvailableadmissioncapacityrequestwlanR13 *bool
	BackhauldlBandwidthrequestwlanR13        *bool
	BackhaululBandwidthrequestwlanR13        *bool
	ChannelutilizationrequestwlanR13         *bool
	StationcountrequestwlanR13               *bool
}
