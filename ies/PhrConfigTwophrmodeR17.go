package ies

import "rrc/utils"

// PHR-Config-twoPHRMode-r17 ::= ENUMERATED
type PhrConfigTwophrmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	PhrConfigTwophrmodeR17EnumeratedNothing = iota
	PhrConfigTwophrmodeR17EnumeratedEnabled
)
