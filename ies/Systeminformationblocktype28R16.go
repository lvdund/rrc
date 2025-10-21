package ies

import "rrc/utils"

// SystemInformationBlockType28-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype28R16 struct {
	SegmentnumberR16         utils.INTEGER
	SegmenttypeR16           utils.ENUMERATED
	SegmentcontainerR16      utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
