package ies

import "rrc/utils"

// SlotOrSubslotPDSCH-Config-r15-setup-tbsIndexAlt-STTI-r15 ::= ENUMERATED
type SlotorsubslotpdschConfigR15SetupTbsindexaltSttiR15 struct {
	Value utils.ENUMERATED
}

const (
	SlotorsubslotpdschConfigR15SetupTbsindexaltSttiR15EnumeratedNothing = iota
	SlotorsubslotpdschConfigR15SetupTbsindexaltSttiR15EnumeratedA33
)
