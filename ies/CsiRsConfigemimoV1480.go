package ies

// CSI-RS-ConfigEMIMO-v1480 ::= CHOICE
const (
	CsiRsConfigemimoV1480ChoiceNothing = iota
	CsiRsConfigemimoV1480ChoiceRelease
	CsiRsConfigemimoV1480ChoiceSetup
)

type CsiRsConfigemimoV1480 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigemimoV1480Setup
}
