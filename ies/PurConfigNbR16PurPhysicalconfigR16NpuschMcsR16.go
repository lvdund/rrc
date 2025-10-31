package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-PhysicalConfig-r16-npusch-MCS-r16 ::= CHOICE
const (
	PurConfigNbR16PurPhysicalconfigR16NpuschMcsR16ChoiceNothing = iota
	PurConfigNbR16PurPhysicalconfigR16NpuschMcsR16ChoiceSingletone
	PurConfigNbR16PurPhysicalconfigR16NpuschMcsR16ChoiceMultitone
)

type PurConfigNbR16PurPhysicalconfigR16NpuschMcsR16 struct {
	Choice     uint64
	Singletone *utils.INTEGER `lb:0,ub:10`
	Multitone  *utils.INTEGER `lb:0,ub:13`
}
