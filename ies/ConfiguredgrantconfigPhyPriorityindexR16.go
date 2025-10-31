package ies

import "rrc/utils"

// ConfiguredGrantConfig-phy-PriorityIndex-r16 ::= ENUMERATED
type ConfiguredgrantconfigPhyPriorityindexR16 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigPhyPriorityindexR16EnumeratedNothing = iota
	ConfiguredgrantconfigPhyPriorityindexR16EnumeratedP0
	ConfiguredgrantconfigPhyPriorityindexR16EnumeratedP1
)
