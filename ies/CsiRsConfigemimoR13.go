package ies

// CSI-RS-ConfigEMIMO-r13 ::= CHOICE
const (
	CsiRsConfigemimoR13ChoiceNothing = iota
	CsiRsConfigemimoR13ChoiceRelease
	CsiRsConfigemimoR13ChoiceSetup
)

type CsiRsConfigemimoR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigemimoR13Setup
}
