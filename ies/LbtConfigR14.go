package ies

import "rrc/utils"

// LBT-Config-r14 ::= CHOICE
const (
	LbtConfigR14ChoiceNothing = iota
	LbtConfigR14ChoiceMaxenergydetectionthresholdR14
	LbtConfigR14ChoiceEnergydetectionthresholdoffsetR14
)

type LbtConfigR14 struct {
	Choice                            uint64
	MaxenergydetectionthresholdR14    *utils.INTEGER `lb:-85,ub:-52`
	EnergydetectionthresholdoffsetR14 *utils.INTEGER `lb:-13,ub:20`
}
