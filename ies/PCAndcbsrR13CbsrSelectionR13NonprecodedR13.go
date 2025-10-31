package ies

import "rrc/utils"

// P-C-AndCBSR-r13-cbsr-Selection-r13-nonPrecoded-r13 ::= SEQUENCE
type PCAndcbsrR13CbsrSelectionR13NonprecodedR13 struct {
	Codebooksubsetrestriction1R13 utils.BITSTRING
	Codebooksubsetrestriction2R13 utils.BITSTRING
}
