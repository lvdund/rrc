package ies

import "rrc/utils"

// MRB-RLC-ConfigBroadcast-r17-sn-FieldLength-r17 ::= ENUMERATED
type MrbRlcConfigbroadcastR17SnFieldlengthR17 struct {
	Value utils.ENUMERATED
}

const (
	MrbRlcConfigbroadcastR17SnFieldlengthR17EnumeratedNothing = iota
	MrbRlcConfigbroadcastR17SnFieldlengthR17EnumeratedSize6
)
