package ies

import "rrc/utils"

// SL-ResourcePool-r16-sl-TimeWindowSizeCR-r16 ::= ENUMERATED
type SlResourcepoolR16SlTimewindowsizecrR16 struct {
	Value utils.ENUMERATED
}

const (
	SlResourcepoolR16SlTimewindowsizecrR16EnumeratedNothing = iota
	SlResourcepoolR16SlTimewindowsizecrR16EnumeratedMs1000
	SlResourcepoolR16SlTimewindowsizecrR16EnumeratedSlot1000
)
