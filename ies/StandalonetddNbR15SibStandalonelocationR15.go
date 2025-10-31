package ies

import "rrc/utils"

// StandaloneTDD-NB-r15-sib-StandaloneLocation-r15 ::= ENUMERATED
type StandalonetddNbR15SibStandalonelocationR15 struct {
	Value utils.ENUMERATED
}

const (
	StandalonetddNbR15SibStandalonelocationR15EnumeratedNothing = iota
	StandalonetddNbR15SibStandalonelocationR15EnumeratedLower
	StandalonetddNbR15SibStandalonelocationR15EnumeratedHigher
)
