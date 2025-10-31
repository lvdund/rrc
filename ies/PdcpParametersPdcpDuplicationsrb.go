package ies

import "rrc/utils"

// PDCP-Parameters-pdcp-DuplicationSRB ::= ENUMERATED
type PdcpParametersPdcpDuplicationsrb struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersPdcpDuplicationsrbEnumeratedNothing = iota
	PdcpParametersPdcpDuplicationsrbEnumeratedSupported
)
