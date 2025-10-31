package ies

import "rrc/utils"

// SIB12-r16 ::= SEQUENCE
type Sib12R16 struct {
	SegmentnumberR16    utils.INTEGER `lb:0,ub:63`
	SegmenttypeR16      Sib12R16SegmenttypeR16
	SegmentcontainerR16 utils.OCTETSTRING
}
