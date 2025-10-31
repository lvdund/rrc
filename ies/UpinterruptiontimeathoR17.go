package ies

import "rrc/utils"

// UPInterruptionTimeAtHO-r17 ::= utils.INTEGER (0..1023)
type UpinterruptiontimeathoR17 struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
