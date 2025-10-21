package ies

import "rrc/utils"

// P-C-AndCBSR-r15 ::= SEQUENCE
type PCAndcbsrR15 struct {
	PCR15                         utils.INTEGER
	Codebooksubsetrestriction4R15 utils.BITSTRING
}
