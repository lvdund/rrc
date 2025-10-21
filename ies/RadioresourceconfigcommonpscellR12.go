package ies

import "rrc/utils"

// RadioResourceConfigCommonPSCell-r12 ::= SEQUENCE
// Extensible
type RadioresourceconfigcommonpscellR12 struct {
	BasicfieldsR12                    RadioresourceconfigcommonscellR10
	PucchConfigcommonR12              PucchConfigcommon
	RachConfigcommonR12               RachConfigcommon
	UplinkpowercontrolcommonpscellR12 UplinkpowercontrolcommonpscellR12
}
