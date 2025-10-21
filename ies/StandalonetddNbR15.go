package ies

import "rrc/utils"

// StandaloneTDD-NB-r15 ::= SEQUENCE
type StandalonetddNbR15 struct {
	SibStandalonelocationR15 utils.ENUMERATED
	Spare                    utils.BITSTRING
}
