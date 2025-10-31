package ies

import "rrc/utils"

// PhysCellId ::= utils.INTEGER (0..503)
type Physcellid struct {
	Value utils.INTEGER `lb:0,ub:503`
}
