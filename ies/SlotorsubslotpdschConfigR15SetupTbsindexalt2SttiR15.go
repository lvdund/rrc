package ies

import "rrc/utils"

// SlotOrSubslotPDSCH-Config-r15-setup-tbsIndexAlt2-STTI-r15 ::= ENUMERATED
type SlotorsubslotpdschConfigR15SetupTbsindexalt2SttiR15 struct {
	Value utils.ENUMERATED
}

const (
	SlotorsubslotpdschConfigR15SetupTbsindexalt2SttiR15EnumeratedNothing = iota
	SlotorsubslotpdschConfigR15SetupTbsindexalt2SttiR15EnumeratedB33
)
