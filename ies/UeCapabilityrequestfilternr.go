package ies

// UE-CapabilityRequestFilterNR ::= SEQUENCE
type UeCapabilityrequestfilternr struct {
	Frequencybandlistfilter *Freqbandlist
	Noncriticalextension    *UeCapabilityrequestfilternrV1540
}
