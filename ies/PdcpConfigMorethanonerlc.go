package ies

import "rrc/utils"

// PDCP-Config-moreThanOneRLC ::= SEQUENCE
type PdcpConfigMorethanonerlc struct {
	Primarypath          *PdcpConfigMorethanonerlcPrimarypath
	UlDatasplitthreshold *UlDatasplitthreshold
	PdcpDuplication      *utils.BOOLEAN
}
