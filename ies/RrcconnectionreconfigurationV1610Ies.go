package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1610-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1610Ies struct {
	ConditionalreconfigurationR16 *ConditionalreconfigurationR16
	DapsSourcereleaseR16          *utils.ENUMERATED
	TdmPatternconfig2R16          *TdmPatternconfigR15
	SlConfigdedicatedfornrR16     *utils.OCTETSTRING
	SlSsbPriorityeutraR16         *utils.INTEGER
	Noncriticalextension          *interface{}
}
