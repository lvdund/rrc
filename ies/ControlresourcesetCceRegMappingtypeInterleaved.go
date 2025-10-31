package ies

import "rrc/utils"

// ControlResourceSet-cce-REG-MappingType-interleaved ::= SEQUENCE
type ControlresourcesetCceRegMappingtypeInterleaved struct {
	RegBundlesize   ControlresourcesetCceRegMappingtypeInterleavedRegBundlesize
	Interleaversize ControlresourcesetCceRegMappingtypeInterleavedInterleaversize
	Shiftindex      *utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
}
