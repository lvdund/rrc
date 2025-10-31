package ies

import "rrc/utils"

// BFR-CSIRS-Resource ::= SEQUENCE
// Extensible
type BfrCsirsResource struct {
	CsiRs           NzpCsiRsResourceid
	RaOccasionlist  *[]utils.INTEGER `lb:1,ub:maxRAOccasionspercsirs`
	RaPreambleindex *utils.INTEGER   `lb:0,ub:63`
}
