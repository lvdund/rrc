package ies

// SCG-Configuration-v13c0 ::= CHOICE
const (
	ScgConfigurationV13c0ChoiceNothing = iota
	ScgConfigurationV13c0ChoiceRelease
	ScgConfigurationV13c0ChoiceSetup
)

type ScgConfigurationV13c0 struct {
	Choice  uint64
	Release *struct{}
	Setup   *ScgConfigurationV13c0Setup
}
