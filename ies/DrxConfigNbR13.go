package ies

// DRX-Config-NB-r13 ::= CHOICE
const (
	DrxConfigNbR13ChoiceNothing = iota
	DrxConfigNbR13ChoiceRelease
	DrxConfigNbR13ChoiceSetup
)

type DrxConfigNbR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *DrxConfigNbR13Setup
}
