package ies

// CSI-RS-Config-r15 ::= CHOICE
const (
	CsiRsConfigR15ChoiceNothing = iota
	CsiRsConfigR15ChoiceRelease
	CsiRsConfigR15ChoiceSetup
)

type CsiRsConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigR15Setup
}
