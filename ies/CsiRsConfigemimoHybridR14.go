package ies

// CSI-RS-ConfigEMIMO-Hybrid-r14 ::= CHOICE
const (
	CsiRsConfigemimoHybridR14ChoiceNothing = iota
	CsiRsConfigemimoHybridR14ChoiceRelease
	CsiRsConfigemimoHybridR14ChoiceSetup
)

type CsiRsConfigemimoHybridR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigemimoHybridR14Setup
}
