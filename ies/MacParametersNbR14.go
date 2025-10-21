package ies

import "rrc/utils"

// MAC-Parameters-NB-r14 ::= SEQUENCE
type MacParametersNbR14 struct {
	DatainactmonR14 *utils.ENUMERATED
	RaiSupportR14   *utils.ENUMERATED
}
