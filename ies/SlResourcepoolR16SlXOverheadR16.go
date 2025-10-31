package ies

import "rrc/utils"

// SL-ResourcePool-r16-sl-X-Overhead-r16 ::= ENUMERATED
type SlResourcepoolR16SlXOverheadR16 struct {
	Value utils.ENUMERATED
}

const (
	SlResourcepoolR16SlXOverheadR16EnumeratedNothing = iota
	SlResourcepoolR16SlXOverheadR16EnumeratedN0
	SlResourcepoolR16SlXOverheadR16EnumeratedN3
	SlResourcepoolR16SlXOverheadR16EnumeratedN6
	SlResourcepoolR16SlXOverheadR16EnumeratedN9
)
