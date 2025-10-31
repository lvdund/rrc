package ies

import "rrc/utils"

// ControlResourceSetZero ::= utils.INTEGER (0..15)
type Controlresourcesetzero struct {
	Value utils.INTEGER `lb:0,ub:15`
}
