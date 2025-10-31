package ies

// CSI-RS-ConfigNZP-EMIMO-r13 ::= CHOICE
const (
	CsiRsConfignzpEmimoR13ChoiceNothing = iota
	CsiRsConfignzpEmimoR13ChoiceRelease
	CsiRsConfignzpEmimoR13ChoiceSetup
)

type CsiRsConfignzpEmimoR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfignzpEmimoR13Setup
}
