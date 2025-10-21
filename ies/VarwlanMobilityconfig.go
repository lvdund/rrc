package ies

import "rrc/utils"

// VarWLAN-MobilityConfig ::= SEQUENCE
type VarwlanMobilityconfig struct {
	WlanMobilitysetR13     *WlanIdListR13
	Successreportrequested *utils.ENUMERATED
	WlanSuspendconfigR14   *WlanSuspendconfigR14
}
