package ies

import "rrc/utils"

// P-C-AndCBSR-r13 ::= SEQUENCE
// Extensible
type PCAndcbsrR13 struct {
	PCR13            utils.INTEGER
	CbsrSelectionR13 interface{}
}
