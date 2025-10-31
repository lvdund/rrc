package ies

import "rrc/utils"

// PowerPrefIndicationConfig-r11 ::= CHOICE
const (
	PowerprefindicationconfigR11ChoiceNothing = iota
	PowerprefindicationconfigR11ChoiceRelease
	PowerprefindicationconfigR11ChoiceSetup
)

type PowerprefindicationconfigR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *utils.ENUMERATED
}
