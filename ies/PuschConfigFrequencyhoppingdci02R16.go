package ies

import "rrc/utils"

// PUSCH-Config-frequencyHoppingDCI-0-2-r16 ::= CHOICE
const (
	PuschConfigFrequencyhoppingdci02R16ChoiceNothing = iota
	PuschConfigFrequencyhoppingdci02R16ChoicePuschReptypea
	PuschConfigFrequencyhoppingdci02R16ChoicePuschReptypeb
)

type PuschConfigFrequencyhoppingdci02R16 struct {
	Choice        uint64
	PuschReptypea *utils.ENUMERATED
	PuschReptypeb *utils.ENUMERATED
}
