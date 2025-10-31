package ies

// CSI-RS-ConfigEMIMO-v1530 ::= CHOICE
const (
	CsiRsConfigemimoV1530ChoiceNothing = iota
	CsiRsConfigemimoV1530ChoiceRelease
	CsiRsConfigemimoV1530ChoiceSetup
)

type CsiRsConfigemimoV1530 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigemimoV1530Setup
}
