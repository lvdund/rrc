package ies

import "rrc/utils"

// PDSCH-Config-xOverheadMulticast-r17 ::= ENUMERATED
type PdschConfigXoverheadmulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigXoverheadmulticastR17EnumeratedNothing = iota
	PdschConfigXoverheadmulticastR17EnumeratedXoh6
	PdschConfigXoverheadmulticastR17EnumeratedXoh12
	PdschConfigXoverheadmulticastR17EnumeratedXoh18
)
