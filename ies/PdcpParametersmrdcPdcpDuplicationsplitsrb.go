package ies

import "rrc/utils"

// PDCP-ParametersMRDC-pdcp-DuplicationSplitSRB ::= ENUMERATED
type PdcpParametersmrdcPdcpDuplicationsplitsrb struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersmrdcPdcpDuplicationsplitsrbEnumeratedNothing = iota
	PdcpParametersmrdcPdcpDuplicationsplitsrbEnumeratedSupported
)
