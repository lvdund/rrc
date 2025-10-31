package ies

// RCLWI-Configuration-r13 ::= CHOICE
const (
	RclwiConfigurationR13ChoiceNothing = iota
	RclwiConfigurationR13ChoiceRelease
	RclwiConfigurationR13ChoiceSetup
)

type RclwiConfigurationR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RclwiConfigurationR13Setup
}
