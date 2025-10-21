package ies

import "rrc/utils"

// SlotOrSubslotPUSCH-Config-r15 ::= CHOICE
// Extensible
type SlotorsubslotpuschConfigR15 interface {
	isSlotorsubslotpuschConfigR15()
}

type SlotorsubslotpuschConfigR15Release struct {
	Value struct{}
}

func (*SlotorsubslotpuschConfigR15Release) isSlotorsubslotpuschConfigR15() {}

type SlotorsubslotpuschConfigR15Setup struct {
	Value interface{}
}

func (*SlotorsubslotpuschConfigR15Setup) isSlotorsubslotpuschConfigR15() {}
