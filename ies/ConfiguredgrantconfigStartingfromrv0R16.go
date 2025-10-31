package ies

import "rrc/utils"

// ConfiguredGrantConfig-startingFromRV0-r16 ::= ENUMERATED
type ConfiguredgrantconfigStartingfromrv0R16 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigStartingfromrv0R16EnumeratedNothing = iota
	ConfiguredgrantconfigStartingfromrv0R16EnumeratedOn
	ConfiguredgrantconfigStartingfromrv0R16EnumeratedOff
)
