package ies

import "rrc/utils"

// SRS-ResourceSet-usage ::= ENUMERATED
type SrsResourcesetUsage struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesetUsageEnumeratedNothing = iota
	SrsResourcesetUsageEnumeratedBeammanagement
	SrsResourcesetUsageEnumeratedCodebook
	SrsResourcesetUsageEnumeratedNoncodebook
	SrsResourcesetUsageEnumeratedAntennaswitching
)
