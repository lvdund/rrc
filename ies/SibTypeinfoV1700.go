package ies

import "rrc/utils"

// SIB-TypeInfo-v1700 ::= SEQUENCE
// Extensible
type SibTypeinfoV1700 struct {
	SibtypeR17   *SibTypeinfoV1700SibtypeR17
	ValuetagR17  *utils.INTEGER `lb:0,ub:31`
	AreascopeR17 *utils.BOOLEAN
}
