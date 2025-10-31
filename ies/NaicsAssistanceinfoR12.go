package ies

// NAICS-AssistanceInfo-r12 ::= CHOICE
const (
	NaicsAssistanceinfoR12ChoiceNothing = iota
	NaicsAssistanceinfoR12ChoiceRelease
	NaicsAssistanceinfoR12ChoiceSetup
)

type NaicsAssistanceinfoR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *NaicsAssistanceinfoR12Setup
}
