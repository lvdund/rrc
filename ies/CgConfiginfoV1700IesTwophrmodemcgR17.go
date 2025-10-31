package ies

import "rrc/utils"

// CG-ConfigInfo-v1700-IEs-twoPHRModeMCG-r17 ::= ENUMERATED
type CgConfiginfoV1700IesTwophrmodemcgR17 struct {
	Value utils.ENUMERATED
}

const (
	CgConfiginfoV1700IesTwophrmodemcgR17EnumeratedNothing = iota
	CgConfiginfoV1700IesTwophrmodemcgR17EnumeratedEnabled
)
