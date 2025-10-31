package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-PhysicalConfig-r16-npusch-SubCarrierSetIndex-r16 ::= CHOICE
const (
	PurConfigNbR16PurPhysicalconfigR16NpuschSubcarriersetindexR16ChoiceNothing = iota
	PurConfigNbR16PurPhysicalconfigR16NpuschSubcarriersetindexR16ChoiceKhz15
	PurConfigNbR16PurPhysicalconfigR16NpuschSubcarriersetindexR16ChoiceKhz3dot75
)

type PurConfigNbR16PurPhysicalconfigR16NpuschSubcarriersetindexR16 struct {
	Choice    uint64
	Khz15     *utils.INTEGER `lb:0,ub:18`
	Khz3dot75 *utils.INTEGER `lb:0,ub:47`
}
