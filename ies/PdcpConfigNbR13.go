package ies

import "rrc/utils"

// PDCP-Config-NB-r13 ::= SEQUENCE
// Extensible
type PdcpConfigNbR13 struct {
	DiscardtimerR13      *utils.ENUMERATED
	HeadercompressionR13 interface{}
}
