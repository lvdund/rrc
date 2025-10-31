package ies

// SCG-Configuration-r12 ::= CHOICE
// Extensible
const (
	ScgConfigurationR12ChoiceNothing = iota
	ScgConfigurationR12ChoiceRelease
	ScgConfigurationR12ChoiceSetup
)

type ScgConfigurationR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *ScgConfigurationR12Setup
}
