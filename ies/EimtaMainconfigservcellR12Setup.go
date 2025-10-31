package ies

import "rrc/utils"

// EIMTA-MainConfigServCell-r12-setup ::= SEQUENCE
type EimtaMainconfigservcellR12Setup struct {
	EimtaUlDlConfigindexR12      utils.INTEGER `lb:0,ub:5`
	EimtaHarqReferenceconfigR12  EimtaMainconfigservcellR12SetupEimtaHarqReferenceconfigR12
	MbsfnSubframeconfiglistV1250 EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250
}
