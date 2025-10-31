package ies

import "rrc/utils"

// PHR-Config-phr-PeriodicTimer ::= ENUMERATED
type PhrConfigPhrPeriodictimer struct {
	Value utils.ENUMERATED
}

const (
	PhrConfigPhrPeriodictimerEnumeratedNothing = iota
	PhrConfigPhrPeriodictimerEnumeratedSf10
	PhrConfigPhrPeriodictimerEnumeratedSf20
	PhrConfigPhrPeriodictimerEnumeratedSf50
	PhrConfigPhrPeriodictimerEnumeratedSf100
	PhrConfigPhrPeriodictimerEnumeratedSf200
	PhrConfigPhrPeriodictimerEnumeratedSf500
	PhrConfigPhrPeriodictimerEnumeratedSf1000
	PhrConfigPhrPeriodictimerEnumeratedInfinity
)
