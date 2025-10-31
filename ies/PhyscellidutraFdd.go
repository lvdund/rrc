package ies

import "rrc/utils"

// PhysCellIdUTRA-FDD ::= utils.INTEGER (0..511)
type PhyscellidutraFdd struct {
	Value utils.INTEGER `lb:0,ub:511`
}
