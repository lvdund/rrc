package ies

import "rrc/utils"

// PH-InfoSCG-twoSRS-PUSCH-Repetition-r17 ::= ENUMERATED
type PhInfoscgTwosrsPuschRepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhInfoscgTwosrsPuschRepetitionR17EnumeratedNothing = iota
	PhInfoscgTwosrsPuschRepetitionR17EnumeratedEnabled
)
