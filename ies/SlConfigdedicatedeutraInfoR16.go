package ies

import "rrc/utils"

// SL-ConfigDedicatedEUTRA-Info-r16 ::= SEQUENCE
type SlConfigdedicatedeutraInfoR16 struct {
	SlConfigdedicatedeutraR16 *utils.OCTETSTRING
	SlTimeoffseteutraListR16  *[]SlTimeoffseteutraR16 `lb:8,ub:8`
}
