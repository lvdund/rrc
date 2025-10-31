package ies

import "rrc/utils"

// SL-ResourcePool-r16-sl-TimeWindowSizeCBR-r16 ::= ENUMERATED
type SlResourcepoolR16SlTimewindowsizecbrR16 struct {
	Value utils.ENUMERATED
}

const (
	SlResourcepoolR16SlTimewindowsizecbrR16EnumeratedNothing = iota
	SlResourcepoolR16SlTimewindowsizecbrR16EnumeratedMs100
	SlResourcepoolR16SlTimewindowsizecbrR16EnumeratedSlot100
)
