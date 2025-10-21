package ies

import "rrc/utils"

// CE-Parameters-v1320 ::= SEQUENCE
type CeParametersV1320 struct {
	Intrafreqa3CeModeaR13 *utils.ENUMERATED
	Intrafreqa3CeModebR13 *utils.ENUMERATED
	IntrafreqhoCeModeaR13 *utils.ENUMERATED
	IntrafreqhoCeModebR13 *utils.ENUMERATED
}
