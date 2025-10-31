package ies

// DRX-Config ::= CHOICE
const (
	DrxConfigChoiceNothing = iota
	DrxConfigChoiceRelease
	DrxConfigChoiceSetup
)

type DrxConfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *DrxConfigSetup
}
