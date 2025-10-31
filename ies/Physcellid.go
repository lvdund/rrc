package ies

import "rrc/utils"

// PhysCellId ::= utils.INTEGER (0..1007)
type Physcellid struct {
	Value utils.INTEGER `lb:0,ub:1007`
}
