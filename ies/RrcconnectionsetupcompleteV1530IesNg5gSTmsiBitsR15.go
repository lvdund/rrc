package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1530-IEs-ng-5G-S-TMSI-Bits-r15 ::= CHOICE
const (
	RrcconnectionsetupcompleteV1530IesNg5gSTmsiBitsR15ChoiceNothing = iota
	RrcconnectionsetupcompleteV1530IesNg5gSTmsiBitsR15ChoiceNg5gSTmsiR15
	RrcconnectionsetupcompleteV1530IesNg5gSTmsiBitsR15ChoiceNg5gSTmsiPart2R15
)

type RrcconnectionsetupcompleteV1530IesNg5gSTmsiBitsR15 struct {
	Choice            uint64
	Ng5gSTmsiR15      *Ng5gSTmsiR15
	Ng5gSTmsiPart2R15 *utils.BITSTRING `lb:8,ub:8`
}
