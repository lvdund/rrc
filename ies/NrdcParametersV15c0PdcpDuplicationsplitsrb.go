package ies

import "rrc/utils"

// NRDC-Parameters-v15c0-pdcp-DuplicationSplitSRB ::= ENUMERATED
type NrdcParametersV15c0PdcpDuplicationsplitsrb struct {
	Value utils.ENUMERATED
}

const (
	NrdcParametersV15c0PdcpDuplicationsplitsrbEnumeratedNothing = iota
	NrdcParametersV15c0PdcpDuplicationsplitsrbEnumeratedSupported
)
