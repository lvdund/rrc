package ies

// CSI-RS-ConfigZP-ApList-r14 ::= CHOICE
const (
	CsiRsConfigzpAplistR14ChoiceNothing = iota
	CsiRsConfigzpAplistR14ChoiceRelease
	CsiRsConfigzpAplistR14ChoiceSetup
)

type CsiRsConfigzpAplistR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *[]CsiRsConfigzpR11 `lb:1,ub:maxCSIRsZpR11`
}
