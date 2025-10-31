package ies

// SL-DiscTxResource-r13-setup ::= CHOICE
const (
	SlDisctxresourceR13SetupChoiceNothing = iota
	SlDisctxresourceR13SetupChoiceScheduledR13
	SlDisctxresourceR13SetupChoiceUeSelectedR13
)

type SlDisctxresourceR13Setup struct {
	Choice        uint64
	ScheduledR13  *SlDisctxconfigscheduledR13
	UeSelectedR13 *SlDisctxpooldedicatedR13
}
