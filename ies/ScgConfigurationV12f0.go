package ies

// SCG-Configuration-v12f0 ::= CHOICE
const (
	ScgConfigurationV12f0ChoiceNothing = iota
	ScgConfigurationV12f0ChoiceRelease
	ScgConfigurationV12f0ChoiceSetup
)

type ScgConfigurationV12f0 struct {
	Choice  uint64
	Release *struct{}
	Setup   *ScgConfigurationV12f0Setup
}
