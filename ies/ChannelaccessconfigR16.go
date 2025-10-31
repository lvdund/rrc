package ies

import "rrc/utils"

// ChannelAccessConfig-r16 ::= SEQUENCE
type ChannelaccessconfigR16 struct {
	EnergydetectionconfigR16       *ChannelaccessconfigR16EnergydetectionconfigR16
	UlTodlCotSharingedThresholdR16 *utils.INTEGER `lb:0,ub:-52`
	AbsenceofanyothertechnologyR16 *utils.BOOLEAN
}
