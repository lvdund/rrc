package ies

import "rrc/utils"

// SIB17-r17 ::= SEQUENCE
type Sib17R17 struct {
	SegmentnumberR17    utils.INTEGER `lb:0,ub:63`
	SegmenttypeR17      Sib17R17SegmenttypeR17
	SegmentcontainerR17 utils.OCTETSTRING
}
