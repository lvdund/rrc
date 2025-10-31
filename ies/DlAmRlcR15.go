package ies

import "rrc/utils"

// DL-AM-RLC-r15 ::= SEQUENCE
type DlAmRlcR15 struct {
	TReorderingR15        TReordering
	TStatusprohibitR15    TStatusprohibit
	ExtendedRlcLiFieldR15 utils.BOOLEAN
}
