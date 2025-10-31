package ies

import "rrc/utils"

// ConfiguredGrantConfig-mcs-Table ::= ENUMERATED
type ConfiguredgrantconfigMcsTable struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigMcsTableEnumeratedNothing = iota
	ConfiguredgrantconfigMcsTableEnumeratedQam256
	ConfiguredgrantconfigMcsTableEnumeratedQam64lowse
)
