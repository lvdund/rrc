package ies

// VarWLAN-MobilityConfig ::= SEQUENCE
type VarwlanMobilityconfig struct {
	WlanMobilitysetR13     *WlanIdListR13
	Successreportrequested *bool
	WlanSuspendconfigR14   *WlanSuspendconfigR14
}
