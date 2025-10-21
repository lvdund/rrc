package ies

import "rrc/utils"

// TMGI-r9 ::= SEQUENCE
type TmgiR9 struct {
	PlmnIdR9    interface{}
	ServiceidR9 utils.OCTETSTRING
}
