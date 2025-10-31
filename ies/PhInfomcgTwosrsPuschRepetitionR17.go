package ies

import "rrc/utils"

// PH-InfoMCG-twoSRS-PUSCH-Repetition-r17 ::= ENUMERATED
type PhInfomcgTwosrsPuschRepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhInfomcgTwosrsPuschRepetitionR17EnumeratedNothing = iota
	PhInfomcgTwosrsPuschRepetitionR17EnumeratedEnabled
)
