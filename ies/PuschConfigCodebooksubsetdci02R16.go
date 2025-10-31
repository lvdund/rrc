package ies

import "rrc/utils"

// PUSCH-Config-codebookSubsetDCI-0-2-r16 ::= ENUMERATED
type PuschConfigCodebooksubsetdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigCodebooksubsetdci02R16EnumeratedNothing = iota
	PuschConfigCodebooksubsetdci02R16EnumeratedFullyandpartialandnoncoherent
	PuschConfigCodebooksubsetdci02R16EnumeratedPartialandnoncoherent
	PuschConfigCodebooksubsetdci02R16EnumeratedNoncoherent
)
