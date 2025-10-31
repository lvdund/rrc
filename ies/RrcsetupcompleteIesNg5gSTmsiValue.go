package ies

import "rrc/utils"

// RRCSetupComplete-IEs-ng-5G-S-TMSI-Value ::= CHOICE
const (
	RrcsetupcompleteIesNg5gSTmsiValueChoiceNothing = iota
	RrcsetupcompleteIesNg5gSTmsiValueChoiceNg5gSTmsi
	RrcsetupcompleteIesNg5gSTmsiValueChoiceNg5gSTmsiPart2
)

type RrcsetupcompleteIesNg5gSTmsiValue struct {
	Choice         uint64
	Ng5gSTmsi      *Ng5gSTmsi
	Ng5gSTmsiPart2 *utils.BITSTRING `lb:9,ub:9`
}
