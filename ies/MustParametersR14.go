package ies

import "rrc/utils"

// MUST-Parameters-r14 ::= SEQUENCE
type MustParametersR14 struct {
	MustTm234Upto2txR14                   *utils.ENUMERATED
	MustTm89UptooneinterferinglayerR14    *utils.ENUMERATED
	MustTm10UptooneinterferinglayerR14    *utils.ENUMERATED
	MustTm89UptothreeinterferinglayersR14 *utils.ENUMERATED
	MustTm10UptothreeinterferinglayersR14 *utils.ENUMERATED
}
