package ies

import "rrc/utils"

// TMGI-r17 ::= SEQUENCE
type TmgiR17 struct {
	PlmnIdR17    TmgiR17PlmnIdR17
	ServiceidR17 utils.OCTETSTRING `lb:3,ub:3`
}
