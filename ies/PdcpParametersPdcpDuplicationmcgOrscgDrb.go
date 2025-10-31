package ies

import "rrc/utils"

// PDCP-Parameters-pdcp-DuplicationMCG-OrSCG-DRB ::= ENUMERATED
type PdcpParametersPdcpDuplicationmcgOrscgDrb struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersPdcpDuplicationmcgOrscgDrbEnumeratedNothing = iota
	PdcpParametersPdcpDuplicationmcgOrscgDrbEnumeratedSupported
)
