package ies

// SL-V2X-ConfigDedicated-r14-commTxResources-r14-setup ::= CHOICE
const (
	SlV2xConfigdedicatedR14CommtxresourcesR14SetupChoiceNothing = iota
	SlV2xConfigdedicatedR14CommtxresourcesR14SetupChoiceScheduledR14
	SlV2xConfigdedicatedR14CommtxresourcesR14SetupChoiceUeSelectedR14
)

type SlV2xConfigdedicatedR14CommtxresourcesR14Setup struct {
	Choice        uint64
	ScheduledR14  *SlV2xConfigdedicatedR14CommtxresourcesR14SetupScheduledR14
	UeSelectedR14 *SlV2xConfigdedicatedR14CommtxresourcesR14SetupUeSelectedR14
}
