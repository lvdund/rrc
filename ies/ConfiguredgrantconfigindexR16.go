package ies

import "rrc/utils"

// ConfiguredGrantConfigIndex-r16 ::= utils.INTEGER (0.. maxNrofConfiguredGrantConfig-1-r16)
type ConfiguredgrantconfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofConfiguredGrantConfig1R16`
}
