package ies

import "rrc/utils"

// EUTRA-PhysCellId ::= utils.INTEGER (0..503)
type EutraPhyscellid struct {
	Value utils.INTEGER `lb:0,ub:503`
}
