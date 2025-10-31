package ies

import "rrc/utils"

// NAICS-Capability-Entry ::= SEQUENCE
// Extensible
type NaicsCapabilityEntry struct {
	NumberofnaicsCapablecc utils.INTEGER `lb:0,ub:5`
	Numberofaggregatedprb  NaicsCapabilityEntryNumberofaggregatedprb
}
