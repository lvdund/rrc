package ies

import "rrc/utils"

// QFI ::= utils.INTEGER (0..maxQFI)
type Qfi struct {
	Value utils.INTEGER `lb:0,ub:maxQFI`
}
