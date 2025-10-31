package ies

import "rrc/utils"

// DummyJ ::= SEQUENCE
type Dummyj struct {
	MaxenergydetectionthresholdR16    utils.INTEGER  `lb:0,ub:-52`
	EnergydetectionthresholdoffsetR16 utils.INTEGER  `lb:0,ub:-13`
	UlTodlCotSharingedThresholdR16    *utils.INTEGER `lb:0,ub:-52`
	AbsenceofanyothertechnologyR16    *utils.BOOLEAN
}
