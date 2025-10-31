package ies

// PDSCH-Config-dl-OrJointTCI-StateList-r17 ::= CHOICE
const (
	PdschConfigDlOrjointtciStatelistR17ChoiceNothing = iota
	PdschConfigDlOrjointtciStatelistR17ChoiceExplicitlist
	PdschConfigDlOrjointtciStatelistR17ChoiceUnifiedtciStaterefR17
)

type PdschConfigDlOrjointtciStatelistR17 struct {
	Choice                uint64
	Explicitlist          *PdschConfigDlOrjointtciStatelistR17Explicitlist
	UnifiedtciStaterefR17 *ServingcellandbwpIdR17
}
