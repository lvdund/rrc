package ies

import "rrc/utils"

// PerRACSI-RSInfo-r16 ::= SEQUENCE
type PerracsiRsinfoR16 struct {
	CsiRsIndexR16                   CsiRsIndex
	NumberofpreamblessentoncsiRsR16 utils.INTEGER `lb:0,ub:200`
}
