package ies

import "rrc/utils"

// PDCP-Config ::= SEQUENCE
// Extensible
type PdcpConfig struct {
	Discardtimer      *utils.ENUMERATED
	RlcAm             *interface{}
	RlcUm             *interface{}
	Headercompression interface{}
}
