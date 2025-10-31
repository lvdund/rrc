package ies

import "rrc/utils"

// DL-PRS-Info-r16 ::= SEQUENCE
type DlPrsInfoR16 struct {
	DlPrsIdR16            utils.INTEGER  `lb:0,ub:255`
	DlPrsResourcesetidR16 utils.INTEGER  `lb:0,ub:7`
	DlPrsResourceidR16    *utils.INTEGER `lb:0,ub:63`
}
