package ies

import "rrc/utils"

// PUSCH-Config-frequencyHoppingDCI-0-1-r16 ::= ENUMERATED
type PuschConfigFrequencyhoppingdci01R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigFrequencyhoppingdci01R16EnumeratedNothing = iota
	PuschConfigFrequencyhoppingdci01R16EnumeratedInterrepetition
	PuschConfigFrequencyhoppingdci01R16EnumeratedInterslot
)
