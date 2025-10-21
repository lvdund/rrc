package ies

import "rrc/utils"

// CE-Parameters-r13 ::= SEQUENCE
type CeParametersR13 struct {
	CeModeaR13 *utils.ENUMERATED
	CeModebR13 *utils.ENUMERATED
}
