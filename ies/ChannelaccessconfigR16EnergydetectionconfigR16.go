package ies

import "rrc/utils"

// ChannelAccessConfig-r16-energyDetectionConfig-r16 ::= CHOICE
const (
	ChannelaccessconfigR16EnergydetectionconfigR16ChoiceNothing = iota
	ChannelaccessconfigR16EnergydetectionconfigR16ChoiceMaxenergydetectionthresholdR16
	ChannelaccessconfigR16EnergydetectionconfigR16ChoiceEnergydetectionthresholdoffsetR16
)

type ChannelaccessconfigR16EnergydetectionconfigR16 struct {
	Choice                            uint64
	MaxenergydetectionthresholdR16    *utils.INTEGER `lb:-85,ub:-52`
	EnergydetectionthresholdoffsetR16 *utils.INTEGER `lb:-13,ub:20`
}
