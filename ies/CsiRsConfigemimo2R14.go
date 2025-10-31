package ies

// CSI-RS-ConfigEMIMO2-r14 ::= CHOICE
const (
	CsiRsConfigemimo2R14ChoiceNothing = iota
	CsiRsConfigemimo2R14ChoiceRelease
	CsiRsConfigemimo2R14ChoiceSetup
)

type CsiRsConfigemimo2R14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigbeamformedR14
}
