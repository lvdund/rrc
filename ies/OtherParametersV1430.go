package ies

import "rrc/utils"

// Other-Parameters-v1430 ::= SEQUENCE
type OtherParametersV1430 struct {
	BwprefindR14        *utils.ENUMERATED
	RlmReportsupportR14 *utils.ENUMERATED
}
