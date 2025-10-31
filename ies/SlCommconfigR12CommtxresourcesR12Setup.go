package ies

// SL-CommConfig-r12-commTxResources-r12-setup ::= CHOICE
const (
	SlCommconfigR12CommtxresourcesR12SetupChoiceNothing = iota
	SlCommconfigR12CommtxresourcesR12SetupChoiceScheduledR12
	SlCommconfigR12CommtxresourcesR12SetupChoiceUeSelectedR12
)

type SlCommconfigR12CommtxresourcesR12Setup struct {
	Choice        uint64
	ScheduledR12  *SlCommconfigR12CommtxresourcesR12SetupScheduledR12
	UeSelectedR12 *SlCommconfigR12CommtxresourcesR12SetupUeSelectedR12
}
