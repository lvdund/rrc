package ies

import "rrc/utils"

// ConfiguredGrantConfigIndexMAC-r16 ::= utils.INTEGER (0.. maxNrofConfiguredGrantConfigMAC-1-r16)
type ConfiguredgrantconfigindexmacR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofConfiguredGrantConfigMAC1R16`
}
