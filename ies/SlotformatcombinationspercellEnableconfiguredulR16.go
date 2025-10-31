package ies

import "rrc/utils"

// SlotFormatCombinationsPerCell-enableConfiguredUL-r16 ::= ENUMERATED
type SlotformatcombinationspercellEnableconfiguredulR16 struct {
	Value utils.ENUMERATED
}

const (
	SlotformatcombinationspercellEnableconfiguredulR16EnumeratedNothing = iota
	SlotformatcombinationspercellEnableconfiguredulR16EnumeratedEnabled
)
