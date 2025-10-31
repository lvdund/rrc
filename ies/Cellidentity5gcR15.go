package ies

import "rrc/utils"

// CellIdentity-5GC-r15 ::= CHOICE
const (
	Cellidentity5gcR15ChoiceNothing = iota
	Cellidentity5gcR15ChoiceCellidentityR15
	Cellidentity5gcR15ChoiceCellidIndexR15
)

type Cellidentity5gcR15 struct {
	Choice          uint64
	CellidentityR15 *Cellidentity
	CellidIndexR15  *utils.INTEGER `lb:1,ub:maxPLMNR11`
}
