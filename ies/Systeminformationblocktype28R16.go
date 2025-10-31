package ies

import "rrc/utils"

// SystemInformationBlockType28-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype28R16 struct {
	SegmentnumberR16         utils.INTEGER `lb:0,ub:63`
	SegmenttypeR16           Systeminformationblocktype28R16SegmenttypeR16
	SegmentcontainerR16      utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
