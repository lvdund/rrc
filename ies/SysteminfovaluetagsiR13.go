package ies

import "rrc/utils"

// SystemInfoValueTagSI-r13 ::= utils.INTEGER (0..3)
type SysteminfovaluetagsiR13 struct {
	Value utils.INTEGER `lb:0,ub:3`
}
