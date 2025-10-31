package ies

// LWA-Configuration-r13 ::= CHOICE
const (
	LwaConfigurationR13ChoiceNothing = iota
	LwaConfigurationR13ChoiceRelease
	LwaConfigurationR13ChoiceSetup
)

type LwaConfigurationR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *LwaConfigurationR13Setup
}
