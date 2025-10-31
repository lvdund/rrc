package ies

// DMRS-Config-r11 ::= CHOICE
const (
	DmrsConfigR11ChoiceNothing = iota
	DmrsConfigR11ChoiceRelease
	DmrsConfigR11ChoiceSetup
)

type DmrsConfigR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *DmrsConfigR11Setup
}
