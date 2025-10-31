package ies

import "rrc/utils"

// PDCP-ParametersMRDC-pdcp-DuplicationSplitDRB ::= ENUMERATED
type PdcpParametersmrdcPdcpDuplicationsplitdrb struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersmrdcPdcpDuplicationsplitdrbEnumeratedNothing = iota
	PdcpParametersmrdcPdcpDuplicationsplitdrbEnumeratedSupported
)
