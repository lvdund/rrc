package ies

import "rrc/utils"

// PhysCellIdUTRA-FDD-r16 ::= utils.INTEGER (0..511)
type PhyscellidutraFddR16 struct {
	Value utils.INTEGER `lb:0,ub:511`
}
