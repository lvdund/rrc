package ies

import "rrc/utils"

// NRDC-Parameters-v15c0-pdcp-DuplicationSplitDRB ::= ENUMERATED
type NrdcParametersV15c0PdcpDuplicationsplitdrb struct {
	Value utils.ENUMERATED
}

const (
	NrdcParametersV15c0PdcpDuplicationsplitdrbEnumeratedNothing = iota
	NrdcParametersV15c0PdcpDuplicationsplitdrbEnumeratedSupported
)
