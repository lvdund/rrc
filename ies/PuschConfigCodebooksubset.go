package ies

import "rrc/utils"

// PUSCH-Config-codebookSubset ::= ENUMERATED
type PuschConfigCodebooksubset struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigCodebooksubsetEnumeratedNothing = iota
	PuschConfigCodebooksubsetEnumeratedFullyandpartialandnoncoherent
	PuschConfigCodebooksubsetEnumeratedPartialandnoncoherent
	PuschConfigCodebooksubsetEnumeratedNoncoherent
)
