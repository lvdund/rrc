package ies

import "rrc/utils"

// MultiDCI-MultiTRP-r16 ::= SEQUENCE
type MultidciMultitrpR16 struct {
	MaxnumbercoresetR16             MultidciMultitrpR16MaxnumbercoresetR16
	MaxnumbercoresetperpoolindexR16 utils.INTEGER `lb:0,ub:3`
	MaxnumberunicastpdschPerpoolR16 MultidciMultitrpR16MaxnumberunicastpdschPerpoolR16
}
