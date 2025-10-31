package ies

import "rrc/utils"

// CFRA-CSIRS-Resource ::= SEQUENCE
// Extensible
type CfraCsirsResource struct {
	CsiRs           CsiRsIndex
	RaOccasionlist  []utils.INTEGER `lb:1,ub:maxRAOccasionspercsirs`
	RaPreambleindex utils.INTEGER   `lb:0,ub:63`
}
