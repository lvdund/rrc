package ies

import "rrc/utils"

// SL-PSFCH-Config-r16-sl-PSFCH-CandidateResourceType-r16 ::= ENUMERATED
type SlPsfchConfigR16SlPsfchCandidateresourcetypeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsfchConfigR16SlPsfchCandidateresourcetypeR16EnumeratedNothing = iota
	SlPsfchConfigR16SlPsfchCandidateresourcetypeR16EnumeratedStartsubch
	SlPsfchConfigR16SlPsfchCandidateresourcetypeR16EnumeratedAllocsubch
)
