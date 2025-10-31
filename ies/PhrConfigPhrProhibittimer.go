package ies

import "rrc/utils"

// PHR-Config-phr-ProhibitTimer ::= ENUMERATED
type PhrConfigPhrProhibittimer struct {
	Value utils.ENUMERATED
}

const (
	PhrConfigPhrProhibittimerEnumeratedNothing = iota
	PhrConfigPhrProhibittimerEnumeratedSf0
	PhrConfigPhrProhibittimerEnumeratedSf10
	PhrConfigPhrProhibittimerEnumeratedSf20
	PhrConfigPhrProhibittimerEnumeratedSf50
	PhrConfigPhrProhibittimerEnumeratedSf100
	PhrConfigPhrProhibittimerEnumeratedSf200
	PhrConfigPhrProhibittimerEnumeratedSf500
	PhrConfigPhrProhibittimerEnumeratedSf1000
)
