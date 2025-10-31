package ies

import "rrc/utils"

// PDSCH-Config-maxNrofCodeWordsScheduledByDCI ::= ENUMERATED
type PdschConfigMaxnrofcodewordsscheduledbydci struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigMaxnrofcodewordsscheduledbydciEnumeratedNothing = iota
	PdschConfigMaxnrofcodewordsscheduledbydciEnumeratedN1
	PdschConfigMaxnrofcodewordsscheduledbydciEnumeratedN2
)
