package ies

import "rrc/utils"

// CE-Parameters-v1370 ::= SEQUENCE
type CeParametersV1370 struct {
	Tm9CeModeaR13 *utils.ENUMERATED
	Tm9CeModebR13 *utils.ENUMERATED
}
