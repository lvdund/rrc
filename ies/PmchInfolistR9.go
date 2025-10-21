package ies

import "rrc/utils"

// PMCH-InfoList-r9 ::= SEQUENCE OF PMCH-Info-r9
// SIZE (0..maxPMCH-PerMBSFN)
type PmchInfolistR9 struct {
	Value utils.Sequence[PmchInfoR9]
}
