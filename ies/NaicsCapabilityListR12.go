package ies

// NAICS-Capability-List-r12 ::= SEQUENCE OF NAICS-Capability-Entry-r12
// SIZE (1..maxNAICS-Entries-r12)
type NaicsCapabilityListR12 struct {
	Value []NaicsCapabilityEntryR12 `lb:1,ub:maxNAICSEntriesR12`
}
