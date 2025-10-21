package ies

import "rrc/utils"

// RadioResourceConfigCommonSIB-NB-r13 ::= SEQUENCE
// Extensible
type RadioresourceconfigcommonsibNbR13 struct {
	RachConfigcommonR13         RachConfigcommonNbR13
	BcchConfigR13               BcchConfigNbR13
	PcchConfigR13               PcchConfigNbR13
	NprachConfigR13             NprachConfigsibNbR13
	NpdschConfigcommonR13       NpdschConfigcommonNbR13
	NpuschConfigcommonR13       NpuschConfigcommonNbR13
	DlGapR13                    *DlGapconfigNbR13
	UplinkpowercontrolcommonR13 UplinkpowercontrolcommonNbR13
}
