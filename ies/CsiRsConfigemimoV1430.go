package ies

// CSI-RS-ConfigEMIMO-v1430 ::= CHOICE
const (
	CsiRsConfigemimoV1430ChoiceNothing = iota
	CsiRsConfigemimoV1430ChoiceRelease
	CsiRsConfigemimoV1430ChoiceSetup
)

type CsiRsConfigemimoV1430 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigemimoV1430Setup
}
