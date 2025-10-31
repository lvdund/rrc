package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-PhysicalConfig-r16 ::= SEQUENCE
type PurConfigNbR16PurPhysicalconfigR16 struct {
	CarrierconfigR16             CarrierconfigdedicatedNbR13
	NpuschNumrusindexR16         utils.INTEGER `lb:0,ub:7`
	NpuschNumrepetitionsindexR16 utils.INTEGER `lb:0,ub:7`
	NpuschSubcarriersetindexR16  PurConfigNbR16PurPhysicalconfigR16NpuschSubcarriersetindexR16
	NpuschMcsR16                 PurConfigNbR16PurPhysicalconfigR16NpuschMcsR16
	P0UeNpuschR16                utils.INTEGER `lb:0,ub:7`
	AlphaR16                     PurConfigNbR16PurPhysicalconfigR16AlphaR16
	NpuschCyclicshiftR16         PurConfigNbR16PurPhysicalconfigR16NpuschCyclicshiftR16
	NpdcchConfigR16              NpdcchConfigdedicatedNbR13
}
