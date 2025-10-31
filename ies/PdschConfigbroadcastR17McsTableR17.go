package ies

import "rrc/utils"

// PDSCH-ConfigBroadcast-r17-mcs-Table-r17 ::= ENUMERATED
type PdschConfigbroadcastR17McsTableR17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigbroadcastR17McsTableR17EnumeratedNothing = iota
	PdschConfigbroadcastR17McsTableR17EnumeratedQam256
	PdschConfigbroadcastR17McsTableR17EnumeratedQam64lowse
)
