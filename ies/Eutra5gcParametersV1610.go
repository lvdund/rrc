package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-v1610 ::= SEQUENCE
type Eutra5gcParametersV1610 struct {
	CeInactivestateR16 *utils.ENUMERATED
	CeEutra5gcR16      *utils.ENUMERATED
}
