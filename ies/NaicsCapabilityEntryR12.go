package ies

import "rrc/utils"

// NAICS-Capability-Entry-r12 ::= SEQUENCE
// Extensible
type NaicsCapabilityEntryR12 struct {
	NumberofnaicsCapableccR12 utils.INTEGER `lb:0,ub:5`
	NumberofaggregatedprbR12  NaicsCapabilityEntryR12NumberofaggregatedprbR12
}
