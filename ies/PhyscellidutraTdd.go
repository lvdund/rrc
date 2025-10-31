package ies

import "rrc/utils"

// PhysCellIdUTRA-TDD ::= utils.INTEGER (0..127)
type PhyscellidutraTdd struct {
	Value utils.INTEGER `lb:0,ub:127`
}
