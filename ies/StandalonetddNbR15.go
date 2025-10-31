package ies

import "rrc/utils"

// StandaloneTDD-NB-r15 ::= SEQUENCE
type StandalonetddNbR15 struct {
	SibStandalonelocationR15 StandalonetddNbR15SibStandalonelocationR15
	Spare                    utils.BITSTRING `lb:5,ub:5`
}
