package ies

import "rrc/utils"

// UE-Capability-NB-r13-multipleDRB-r13 ::= ENUMERATED
type UeCapabilityNbR13MultipledrbR13 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbR13MultipledrbR13EnumeratedNothing = iota
	UeCapabilityNbR13MultipledrbR13EnumeratedSupported
)
