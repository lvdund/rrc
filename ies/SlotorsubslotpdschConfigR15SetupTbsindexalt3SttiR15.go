package ies

import "rrc/utils"

// SlotOrSubslotPDSCH-Config-r15-setup-tbsIndexAlt3-STTI-r15 ::= ENUMERATED
type SlotorsubslotpdschConfigR15SetupTbsindexalt3SttiR15 struct {
	Value utils.ENUMERATED
}

const (
	SlotorsubslotpdschConfigR15SetupTbsindexalt3SttiR15EnumeratedNothing = iota
	SlotorsubslotpdschConfigR15SetupTbsindexalt3SttiR15EnumeratedA37
)
