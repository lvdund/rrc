package ies

// SL-DiscConfig-r12-discTxResources-r12-setup ::= CHOICE
const (
	SlDiscconfigR12DisctxresourcesR12SetupChoiceNothing = iota
	SlDiscconfigR12DisctxresourcesR12SetupChoiceScheduledR12
	SlDiscconfigR12DisctxresourcesR12SetupChoiceUeSelectedR12
)

type SlDiscconfigR12DisctxresourcesR12Setup struct {
	Choice        uint64
	ScheduledR12  *SlDiscconfigR12DisctxresourcesR12SetupScheduledR12
	UeSelectedR12 *SlDiscconfigR12DisctxresourcesR12SetupUeSelectedR12
}
