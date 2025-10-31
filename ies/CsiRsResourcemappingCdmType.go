package ies

import "rrc/utils"

// CSI-RS-ResourceMapping-cdm-Type ::= ENUMERATED
type CsiRsResourcemappingCdmType struct {
	Value utils.ENUMERATED
}

const (
	CsiRsResourcemappingCdmTypeEnumeratedNothing = iota
	CsiRsResourcemappingCdmTypeEnumeratedNocdm
	CsiRsResourcemappingCdmTypeEnumeratedFd_Cdm2
	CsiRsResourcemappingCdmTypeEnumeratedCdm4_Fd2_Td2
	CsiRsResourcemappingCdmTypeEnumeratedCdm8_Fd2_Td4
)
