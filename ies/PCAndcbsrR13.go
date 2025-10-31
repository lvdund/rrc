package ies

import "rrc/utils"

// P-C-AndCBSR-r13 ::= SEQUENCE
// Extensible
type PCAndcbsrR13 struct {
	PCR13            utils.INTEGER `lb:0,ub:15`
	CbsrSelectionR13 PCAndcbsrR13CbsrSelectionR13
}
