package ies

import "rrc/utils"

// PDSCH-ConfigBroadcast-r17-xOverhead-r17 ::= ENUMERATED
type PdschConfigbroadcastR17XoverheadR17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigbroadcastR17XoverheadR17EnumeratedNothing = iota
	PdschConfigbroadcastR17XoverheadR17EnumeratedXoh6
	PdschConfigbroadcastR17XoverheadR17EnumeratedXoh12
	PdschConfigbroadcastR17XoverheadR17EnumeratedXoh18
)
