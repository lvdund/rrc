package ies

import "rrc/utils"

// CG-Config-v1700-IEs-twoPHRModeSCG-r17 ::= ENUMERATED
type CgConfigV1700IesTwophrmodescgR17 struct {
	Value utils.ENUMERATED
}

const (
	CgConfigV1700IesTwophrmodescgR17EnumeratedNothing = iota
	CgConfigV1700IesTwophrmodescgR17EnumeratedEnabled
)
