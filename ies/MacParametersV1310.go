package ies

import "rrc/utils"

// MAC-Parameters-v1310 ::= SEQUENCE
type MacParametersV1310 struct {
	ExtendedmacLengthfieldR13 *utils.ENUMERATED
	ExtendedlongdrxR13        *utils.ENUMERATED
}
