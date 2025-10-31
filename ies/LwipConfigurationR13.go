package ies

// LWIP-Configuration-r13 ::= CHOICE
const (
	LwipConfigurationR13ChoiceNothing = iota
	LwipConfigurationR13ChoiceRelease
	LwipConfigurationR13ChoiceSetup
)

type LwipConfigurationR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *LwipConfigurationR13Setup
}
