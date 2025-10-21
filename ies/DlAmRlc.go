package ies

import "rrc/utils"

// DL-AM-RLC ::= SEQUENCE
type DlAmRlc struct {
	TReordering     TReordering
	TStatusprohibit TStatusprohibit
}
