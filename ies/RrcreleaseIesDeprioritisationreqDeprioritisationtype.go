package ies

import "rrc/utils"

// RRCRelease-IEs-deprioritisationReq-deprioritisationType ::= ENUMERATED
type RrcreleaseIesDeprioritisationreqDeprioritisationtype struct {
	Value utils.ENUMERATED
}

const (
	RrcreleaseIesDeprioritisationreqDeprioritisationtypeEnumeratedNothing = iota
	RrcreleaseIesDeprioritisationreqDeprioritisationtypeEnumeratedFrequency
	RrcreleaseIesDeprioritisationreqDeprioritisationtypeEnumeratedNr
)
