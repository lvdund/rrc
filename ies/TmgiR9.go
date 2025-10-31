package ies

import "rrc/utils"

// TMGI-r9 ::= SEQUENCE
type TmgiR9 struct {
	PlmnIdR9    TmgiR9PlmnIdR9
	ServiceidR9 utils.OCTETSTRING `lb:3,ub:3`
}
