package ies

import "rrc/utils"

// SlotOrSubslotPDSCH-Config-r15-setup-resourceAllocation-r15 ::= ENUMERATED
type SlotorsubslotpdschConfigR15SetupResourceallocationR15 struct {
	Value utils.ENUMERATED
}

const (
	SlotorsubslotpdschConfigR15SetupResourceallocationR15EnumeratedNothing = iota
	SlotorsubslotpdschConfigR15SetupResourceallocationR15EnumeratedResourceallocationtype0
	SlotorsubslotpdschConfigR15SetupResourceallocationR15EnumeratedResourceallocationtype2
)
