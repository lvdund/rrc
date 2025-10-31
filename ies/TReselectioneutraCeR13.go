package ies

import "rrc/utils"

// T-ReselectionEUTRA-CE-r13 ::= utils.INTEGER (0..15)
type TReselectioneutraCeR13 struct {
	Value utils.INTEGER `lb:0,ub:15`
}
